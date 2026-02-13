package handlers

import chatv1 "github.com/BeInBloom/grpc-chat/gen/go/chat/v1"

type chatService interface {
	// TODO: define service interface as methods are implemented
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
	return nil
}
