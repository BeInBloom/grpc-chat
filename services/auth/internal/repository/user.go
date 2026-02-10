package repository

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/BeInBloom/grpc-chat/services/auth/internal/models"
)

type UserRepository struct {
	users map[string]*models.User
	mu    sync.RWMutex
}

func New(config models.UserRepositoryConfig) *UserRepository {
	return &UserRepository{
		users: make(map[string]*models.User),
	}
}

func (r *UserRepository) Create(ctx context.Context, user models.User) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = uuid.New().String()
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	r.users[user.ID] = &user

	return user.ID, nil
}

func (r *UserRepository) Get(ctx context.Context, uuid string) (models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[uuid]
	if !ok {
		return models.User{}, ErrUserNotFound
	}

	return *user, nil
}

func (r *UserRepository) Update(ctx context.Context, user models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.users[user.ID]
	if !ok {
		return ErrUserNotFound
	}

	user.UpdatedAt = time.Now()
	r.users[user.ID] = &user

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[id]; !ok {
		return ErrUserNotFound
	}

	delete(r.users, id)

	return nil
}
