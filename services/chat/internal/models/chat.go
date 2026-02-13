package models

import (
	"time"

	"github.com/google/uuid"
)

type ChatType int32

const (
	ChatTypeUnspecified ChatType = iota
	ChatTypeDirect
	ChatTypeGroup
)

type Chat struct {
	ID        uuid.UUID
	Name      string
	Type      ChatType
	CreatedAt time.Time
	UpdatedAt time.Time
}
