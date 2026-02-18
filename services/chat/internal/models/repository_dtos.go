package models

import (
	"time"

	"github.com/google/uuid"
)

type EventDTO struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ChatID    *uuid.UUID
	Type      EventType
	Payload   []byte
	CreatedAt time.Time
}

type ChatDTO struct {
	ID        uuid.UUID
	Name      string
	Type      ChatType
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ChatMemberDTO struct {
	ChatID   uuid.UUID
	UserID   uuid.UUID
	Role     MemberRole
	JoinedAt time.Time
}

type MessageDTO struct {
	ID        uuid.UUID
	ChatID    uuid.UUID
	SenderID  uuid.UUID
	Content   []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

type IdempotencyResult struct {
	MessageID uuid.UUID
	CreatedAt time.Time
}

type SubscriptionInfo struct {
	UserID     uuid.UUID
	LastSeenID uuid.UUID
	CreatedAt  time.Time
}
