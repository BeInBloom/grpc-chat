package models

import (
	"time"

	"github.com/google/uuid"
)

type EventPayloadType string

const (
	PayloadTypeMessageNew         EventPayloadType = "MESSAGE_NEW"
	PayloadTypeMessageUpdated     EventPayloadType = "MESSAGE_UPDATED"
	PayloadTypeMessageDeleted     EventPayloadType = "MESSAGE_DELETED"
	PayloadTypeTypingIndicator    EventPayloadType = "TYPING_INDICATOR"
	PayloadTypeReadReceipt        EventPayloadType = "READ_RECEIPT"
	PayloadTypeSystemNotification EventPayloadType = "SYSTEM_NOTIFICATION"
)

type MessageNewPayload struct {
	MessageID uuid.UUID
	ChatID    uuid.UUID
	SenderID  uuid.UUID
	Content   MessageContent
	CreatedAt time.Time
}

type MessageUpdatedPayload struct {
	MessageID  uuid.UUID
	ChatID     uuid.UUID
	NewContent MessageContent
	UpdatedAt  time.Time
}

type MessageDeletedPayload struct {
	MessageID uuid.UUID
	ChatID    uuid.UUID
	DeletedAt time.Time
}

type TypingIndicatorPayload struct {
	ChatID   uuid.UUID
	UserID   uuid.UUID
	IsTyping bool
}

type ReadReceiptPayload struct {
	ChatID    uuid.UUID
	UserID    uuid.UUID
	MessageID uuid.UUID
	ReadAt    time.Time
}

type SystemNotificationPayload struct {
	Text  string
	Level SystemNotificationLevel
}

func GetPayloadType(eventType EventType) EventPayloadType {
	switch eventType {
	case EventTypeMessageNew:
		return PayloadTypeMessageNew
	case EventTypeMessageUpdated:
		return PayloadTypeMessageUpdated
	case EventTypeMessageDeleted:
		return PayloadTypeMessageDeleted
	case EventTypeTyping:
		return PayloadTypeTypingIndicator
	case EventTypeReadReceipt:
		return PayloadTypeReadReceipt
	case EventTypeSystem:
		return PayloadTypeSystemNotification
	default:
		return ""
	}
}
