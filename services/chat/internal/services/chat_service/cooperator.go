package chatservice

import (
	"bytes"
	"context"
	"errors"
	"sync"

	"github.com/BeInBloom/grpc-chat/services/chat/internal/models"
	"github.com/google/uuid"
)

var (
	ErrBufferOverflow    = errors.New("live channel buffer overflow: client consumption is too slow")
	ErrBrokerSubClosed   = errors.New("global broker subscription channel closed")
	ErrLiveChannelClosed = errors.New("live channel closed unexpectedly")
)

type StreamCooperator struct {
	liveChan   chan models.Event
	cancel     context.CancelCauseFunc
	wg         sync.WaitGroup
	isLiveMode bool
}

func NewStreamCooperator(parentCtx context.Context, bufferSize int) (context.Context, *StreamCooperator) {
	ctx, cancel := context.WithCancelCause(parentCtx)

	return ctx, &StreamCooperator{
		liveChan:   make(chan models.Event, bufferSize),
		cancel:     cancel,
		isLiveMode: false,
	}
}

func (c *StreamCooperator) Close(cause error) {
	c.cancel(cause)
	c.wg.Wait()
}

func (c *StreamCooperator) StartBackgroundProducer(ctx context.Context, subChan <-chan models.Event, unsubscribe func()) {
	c.wg.Go(func() {
		defer unsubscribe()

		for {
			select {
			case <-ctx.Done():
				return

			case event, ok := <-subChan:
				if !ok {
					c.cancel(ErrBrokerSubClosed)
					return
				}

				select {
				case c.liveChan <- event:
				default:
					c.cancel(ErrBufferOverflow)
					return
				}
			}
		}
	})
}

func (c *StreamCooperator) ServeStream(
	ctx context.Context,
	fetchHistory func(ctx context.Context) ([]models.Event, error),
	send func(models.Event) error,
) error {
	history, err := fetchHistory(ctx)
	if err != nil {
		return err
	}

	var cursor uuid.UUID
	if len(history) > 0 {
		cursor = history[len(history)-1].ID

		for _, ev := range history {
			if err := send(ev); err != nil {
				return err
			}
		}
	}

	for {
		select {
		case <-ctx.Done():
			return context.Cause(ctx)

		case ev, ok := <-c.liveChan:
			if !ok {
				return ErrLiveChannelClosed
			}

			if !c.isLiveMode {
				if !IsAfter(ev.ID, cursor) {
					continue
				}

				c.isLiveMode = true
			}

			if err := send(ev); err != nil {
				return err
			}
		}
	}
}

func CompareUUIDv7(a, b uuid.UUID) int {
	return bytes.Compare(a[:], b[:])
}

func IsAfter(eventID, cursorID uuid.UUID) bool {
	return CompareUUIDv7(eventID, cursorID) > 0
}
