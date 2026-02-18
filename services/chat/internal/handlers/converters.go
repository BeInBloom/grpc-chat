package handlers

import (
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	chatv1 "github.com/BeInBloom/grpc-chat/gen/go/chat/v1"
	"github.com/BeInBloom/grpc-chat/services/chat/internal/models"
)

//nolint:unused // will be used when chat service RPCs are implemented
func toMessageContent(c *chatv1.MessageContent) (models.MessageContent, error) {
	if c == nil {
		return models.MessageContent{}, status.Error(codes.InvalidArgument, "content is required")
	}

	mc := models.MessageContent{}

	switch v := c.GetType().(type) {
	case *chatv1.MessageContent_Text:
		mc.Type = models.ContentTypeText
		mc.Ciphertext = v.Text.GetCiphertext()
	default:
		return models.MessageContent{}, status.Error(codes.InvalidArgument, "unsupported content type")
	}

	if c.ReplyToMessageId != nil {
		replyID, err := uuid.Parse(*c.ReplyToMessageId)
		if err != nil {
			return models.MessageContent{}, status.Errorf(codes.InvalidArgument, "invalid reply_to_message_id: %s", err)
		}
		mc.ReplyToMessageID = &replyID
	}

	return mc, nil
}

//nolint:unused // will be used when chat service RPCs are implemented
func toChatID(id string) (uuid.UUID, error) {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, status.Errorf(codes.InvalidArgument, "invalid chat_id: %s", err)
	}

	return parsed, nil
}

//nolint:unused // will be used when chat service RPCs are implemented
func toMessageID(id string) (uuid.UUID, error) {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, status.Errorf(codes.InvalidArgument, "invalid message_id: %s", err)
	}

	return parsed, nil
}

//nolint:unused // will be used when chat service RPCs are implemented
func toProtoMessage(m models.Message) *chatv1.Message {
	return &chatv1.Message{
		Id:        m.ID.String(),
		ChatId:    m.ChatID.String(),
		Sender:    toProtoUser(m.SenderID),
		Content:   toProtoContent(m.Content),
		CreatedAt: timestamppb.New(m.CreatedAt),
		UpdatedAt: timestamppb.New(m.UpdatedAt),
	}
}

//nolint:unused // will be used when chat service RPCs are implemented
func toProtoUser(userID uuid.UUID) *chatv1.User {
	return &chatv1.User{
		Id: userID.String(),
	}
}

//nolint:unused // will be used when chat service RPCs are implemented
func toProtoContent(c models.MessageContent) *chatv1.MessageContent {
	mc := &chatv1.MessageContent{}

	if c.Type == models.ContentTypeText {
		mc.Type = &chatv1.MessageContent_Text{
			Text: &chatv1.TextContent{
				Ciphertext: c.Ciphertext,
			},
		}
	}

	if c.ReplyToMessageID != nil {
		replyStr := c.ReplyToMessageID.String()
		mc.ReplyToMessageId = &replyStr
	}

	return mc
}

//nolint:unused // will be used when chat service RPCs are implemented
func toProtoMessages(messages []models.Message) []*chatv1.Message {
	result := make([]*chatv1.Message, 0, len(messages))
	for _, m := range messages {
		result = append(result, toProtoMessage(m))
	}

	return result
}

func toProtoEvent(e models.Event) *chatv1.ConnectResponse {
	resp := &chatv1.ConnectResponse{
		Id: e.ID.String(),
	}

	switch e.Type {
	case models.EventTypeMessageNew:
		msg, ok := e.Payload.(models.Message)
		if ok {
			resp.Payload = &chatv1.ConnectResponse_MessageNew{
				MessageNew: &chatv1.MessageNew{
					Message: toProtoMessage(msg),
				},
			}
		}
	}

	return resp
}

func toConnectRequest(userUI uuid.UUID, lastEventID uuid.UUID) models.SubscribeRequest {
	return models.SubscribeRequest{
		UserID:      userUI,
		LastEventID: lastEventID,
	}
}
