package services

import (
	"context"

	"github.com/go-playground/validator/v10"

	"github.com/BeInBloom/grpc-chat/services/auth/internal/models"
)

//go:generate mockgen -source=user.go -destination=mocks/mock_repository.go -package=mocks

var validate = validator.New()

type userRepository interface {
	Create(ctx context.Context, user models.User) (string, error)
	Get(ctx context.Context, uuid string) (models.User, error)
	Update(ctx context.Context, user models.User) error
	Delete(ctx context.Context, id string) error
}

type UserService struct {
	repo userRepository
}

func New(repo userRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, user models.User) (string, error) {
	if err := validate.Struct(user); err != nil {
		return "", err
	}

	return s.repo.Create(ctx, user)
}

func (s *UserService) Get(ctx context.Context, id string) (models.User, error) {
	return s.repo.Get(ctx, id)
}

func (s *UserService) Update(ctx context.Context, user models.User) error {
	if err := validate.Struct(user); err != nil {
		return err
	}

	return s.repo.Update(ctx, user)
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
