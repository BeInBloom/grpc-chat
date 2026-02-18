package models

import (
	"time"

	"github.com/google/uuid"
)

type SendMessageRequest struct {
	ChatID         uuid.UUID
	IdempotencyKey string
	Content        MessageContent
}

type SendMessageResponse struct {
	MessageID uuid.UUID
	CreatedAt time.Time
}

type GetHistoryRequest struct {
	ChatID   uuid.UUID
	PageSize int32
	Cursor   string
}

type GetHistoryResponse struct {
	Messages   []Message
	NextCursor string
}

type CreateChatRequest struct {
	Name      string
	Type      ChatType
	MemberIDs []uuid.UUID
}

type CreateChatResponse struct {
	Chat Chat
}

type GetChatRequest struct {
	ChatID uuid.UUID
}

type GetChatResponse struct {
	Chat Chat
}

type ListChatsRequest struct {
	UserID   uuid.UUID
	PageSize int32
	Cursor   string
}

type ListChatsResponse struct {
	Chats      []Chat
	NextCursor string
}

type ChatPreview struct {
	ID          uuid.UUID
	Name        string
	Type        ChatType
	LastMessage *Message
	UnreadCount int32
	UpdatedAt   time.Time
}

type SubscribeRequest struct {
	UserID      uuid.UUID
	LastEventID uuid.UUID
}
