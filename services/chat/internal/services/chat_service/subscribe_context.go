package chatservice

import (
	"context"
	"log/slog"

	"github.com/BeInBloom/grpc-chat/services/chat/internal/models"
)

type subscribeContext struct {
	ctx      context.Context
	req      models.SubscribeRequest
	channels *subscriptionChannels
	service  *ChatService
	cancel   context.CancelFunc
}

func (sc *subscribeContext) replayHistorical() {
	historicalEvents, err := sc.service.eventStore.GetUserEvents(sc.ctx, sc.req.UserID, sc.req.LastEventID, historicalBufferSize)
	if err != nil {
		sc.channels.err <- err
		return
	}
	for _, event := range historicalEvents {
		select {
		case sc.channels.historical <- event:
		case <-sc.ctx.Done():
			return
		}
	}
}

func (sc *subscribeContext) subscribeLive() (<-chan models.Event, error) {
	select {
	case <-sc.ctx.Done():
		return nil, sc.ctx.Err()
	default:
	}
	liveChan, err := sc.service.publisher.Subscribe(sc.ctx, sc.req.UserID)
	if err != nil {
		return nil, err
	}
	go func() {
		defer sc.service.publisher.Unsubscribe(sc.req.UserID)
		for {
			select {
			case event, ok := <-liveChan:
				if !ok {
					return
				}
				select {
				case sc.channels.out <- event:
				case <-sc.ctx.Done():
					return
				}
			case <-sc.ctx.Done():
				return
			}
		}
	}()
	return liveChan, nil
}

func (sc *subscribeContext) fanInEvents(live <-chan models.Event) {
	defer close(sc.channels.out)
	historicalClosed := false
	liveClosed := false
	for {
		select {
		case event, ok := <-sc.channels.historical:
			if ok {
				sc.channels.out <- event
			} else {
				historicalClosed = true
			}
		case event, ok := <-live:
			if ok {
				sc.channels.out <- event
			} else {
				liveClosed = true
			}
		case err := <-sc.channels.err:
			slog.Error("subscription error", "error", err)
			sc.channels.closeAll()
			return
		case <-sc.ctx.Done():
			return
		}
		if historicalClosed && liveClosed {
			return
		}
	}
}

func (sc *subscribeContext) execute() (<-chan models.Event, error) {
	go sc.replayHistorical()
	liveChan, err := sc.subscribeLive()
	if err != nil {
		sc.channels.closeAll()
		return nil, err
	}
	go sc.fanInEvents(liveChan)
	return sc.channels.out, nil
}
