package chatservice

import (
	"context"

	"github.com/BeInBloom/grpc-chat/services/chat/internal/models"
	"github.com/google/uuid"
)

const (
	lastMessages         = 20
	historicalBufferSize = 20
	liveBufferSize       = 100
	outputBufferSize     = 500
	errorBuffer          = 2
)

type (
	eventStore interface {
		GetUserEvents(ctx context.Context, userID uuid.UUID, lastEventID uuid.UUID, limit int32) ([]models.Event, error)
	}

	snapshotter interface{}

	eventPublisher interface {
		Subscribe(ctx context.Context, userID uuid.UUID) (<-chan models.Event, error)
		Unsubscribe(userID uuid.UUID) error
	}

	readModelRepos interface{}
)

type ChatService struct {
	eventStore  eventStore
	snapshotter snapshotter
	publisher   eventPublisher
	readModel   readModelRepos
}

func (s *ChatService) Subscribe(
	ctx context.Context,
	req models.SubscribeRequest,
) (<-chan models.Event, error) {
	ctx, cancel := context.WithCancel(ctx)
	channels := &subscriptionChannels{
		historical: make(chan models.Event, historicalBufferSize),
		out:        make(chan models.Event, outputBufferSize),
		err:        make(chan error, errorBuffer),
	}

	subCtx := &subscribeContext{
		ctx:      ctx,
		req:      req,
		channels: channels,
		service:  s,
		cancel:   cancel,
	}

	return subCtx.execute()
}
