package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	User struct {
		ID        uuid.UUID `validate:"-"`
		Name      string    `validate:"required"`
		Email     string    `validate:"required,email"`
		Password  string    `validate:"required"`
		Role      int32     `validate:"-"`
		CreatedAt time.Time `validate:"-"`
		UpdatedAt time.Time `validate:"-"`
	}
)
