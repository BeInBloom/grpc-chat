package models

type (
	User struct {
		ID    string `validate:"-"`
		Name  string `validate:"required"`
		Email string `validate:"required,email"`
	}
)
