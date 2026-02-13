package models

import (
	"time"

	"github.com/google/uuid"
)

type MemberRole int32

const (
	MemberRoleUnspecified MemberRole = iota
	MemberRoleMember
	MemberRoleAdmin
	MemberRoleOwner
)

type ChatMember struct {
	ChatID   uuid.UUID
	UserID   uuid.UUID
	Role     MemberRole
	JoinedAt time.Time
}
