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

type Message struct {
	ID        uuid.UUID
	ChatID    uuid.UUID
	SenderID  uuid.UUID
	Content   MessageContent
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MessageContent struct {
	Type             ContentType
	Ciphertext       []byte
	ReplyToMessageID *uuid.UUID
}
