package models

import "time"

type (
	User struct {
		ID        string    `validate:"-"`
		Name      string    `validate:"required"`
		Email     string    `validate:"required,email"`
		Password  string    `validate:"required"`
		Role      int32     `validate:"-"`
		CreatedAt time.Time `validate:"-"`
		UpdatedAt time.Time `validate:"-"`
	}
)
