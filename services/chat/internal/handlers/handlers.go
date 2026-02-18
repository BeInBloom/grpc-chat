package handlers

import (
	"context"

	chatv1 "github.com/BeInBloom/grpc-chat/gen/go/chat/v1"
	"github.com/BeInBloom/grpc-chat/services/chat/internal/interceptors"
	"github.com/BeInBloom/grpc-chat/services/chat/internal/models"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type chatService interface {
	Subscribe(ctx context.Context, req models.SubscribeRequest) (<-chan models.Event, error)
	SendMessage(ctx context.Context, req models.SendMessageRequest) (models.SendMessageResponse, error)
	GetHistory(ctx context.Context, req models.GetHistoryRequest) (models.GetHistoryResponse, error)
	CreateChat(ctx context.Context, req models.CreateChatRequest) (models.CreateChatResponse, error)
	GetChat(ctx context.Context, chatID models.GetChatRequest) (models.GetChatResponse, error)
	ListChat(ctx context.Context, req models.ListChatsRequest) (models.ListChatsResponse, error)
}

type Handlers struct {
	chatv1.UnimplementedChatServiceServer
	service chatService
}

func New(service chatService) *Handlers {
	return &Handlers{service: service}
}

func (h *Handlers) Connect(
	req *chatv1.ConnectRequest,
	stream chatv1.ChatService_ConnectServer,
) error {
	ctx := stream.Context()

	userID := interceptors.UserIDFromContext(ctx)

	var lastEventID uuid.UUID
	var err error
	if req.LastEventId != nil {
		lastEventID, err = uuid.Parse(*req.LastEventId)
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "failed to parse last ID event: %v", err)
		}
	}

	eventChan, err := h.service.Subscribe(ctx, toConnectRequest(userID, lastEventID))
	if err != nil {
		return status.Errorf(codes.Internal, "failed to subscribe: %v", err)
	}

	for {
		select {
		case <-ctx.Done():
			return nil

		case event, ok := <-eventChan:
			if !ok {
				return nil
			}

			resp := toProtoEvent(event)

			if err := stream.Send(resp); err != nil {
				return status.Errorf(codes.Unavailable, "failed to send event: %v", err)
			}
		}
	}
}
