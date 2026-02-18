package models

import (
	"time"

	"github.com/google/uuid"
)

type ContentType int32

const (
	ContentTypeUnspecified ContentType = iota
	ContentTypeText
	// ContentTypeImage  // будущее расширение
	// ContentTypeFile
	// ContentTypeVoice
)

type SystemNotificationLevel int32

const (
	SystemNotificationLevelUnspecified SystemNotificationLevel = 0
	SystemNotificationLevelInfo        SystemNotificationLevel = 1
	SystemNotificationLevelWarning     SystemNotificationLevel = 2
	SystemNotificationLevelError       SystemNotificationLevel = 3
)

type EventType string

const (
	EventTypeMessageNew     EventType = "MESSAGE_NEW"
	EventTypeMessageUpdated EventType = "MESSAGE_UPDATED"
	EventTypeMessageDeleted EventType = "MESSAGE_DELETED"
	EventTypeTyping         EventType = "TYPING"
	EventTypeReadReceipt    EventType = "READ_RECEIPT"
	EventTypeSystem         EventType = "SYSTEM"
)

type Message struct {
	ID        uuid.UUID
	ChatID    uuid.UUID
	SenderID  uuid.UUID
	Content   MessageContent
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Event struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Type      EventType
	Payload   any
	CreatedAt time.Time
}

type MessageContent struct {
	Type             ContentType
	Ciphertext       []byte
	ReplyToMessageID *uuid.UUID
}
