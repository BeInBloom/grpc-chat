package interceptors

import (
	"context"

	"github.com/google/uuid"
)

const (
	userID = "user_id"
)

func UserIDFromContext(ctx context.Context) uuid.UUID {
	userID, ok := ctx.Value(userID).(uuid.UUID)
	if ok {
		return userID
	}

	return uuid.UUID{}
}
