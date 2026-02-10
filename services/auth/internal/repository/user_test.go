package repository

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/BeInBloom/grpc-chat/services/auth/internal/models"
)

func newTestRepo() *UserRepository {
	return New(models.UserRepositoryConfig{})
}

func TestUserRepository_Create(t *testing.T) {
	repo := newTestRepo()
	ctx := context.Background()

	userID, err := repo.Create(ctx, models.User{
		Name:     "test",
		Email:    "test@example.com",
		Password: "secret123",
	})

	require.NoError(t, err)
	_, err = uuid.Parse(userID)
	assert.NoError(t, err, "expected valid UUID")

	user, err := repo.Get(ctx, userID)
	require.NoError(t, err)
	assert.False(t, user.CreatedAt.IsZero(), "expected CreatedAt to be set")
	assert.False(t, user.UpdatedAt.IsZero(), "expected UpdatedAt to be set")
}

func TestUserRepository_Get(t *testing.T) {
	repo := newTestRepo()
	ctx := context.Background()

	userID, err := repo.Create(ctx, models.User{
		Name:     "test",
		Email:    "test@example.com",
		Password: "secret123",
	})
	require.NoError(t, err)

	user, err := repo.Get(ctx, userID)

	require.NoError(t, err)
	assert.Equal(t, userID, user.ID)
	assert.Equal(t, "test", user.Name)
	assert.Equal(t, "test@example.com", user.Email)
	assert.Equal(t, "secret123", user.Password)
}

func TestUserRepository_GetNotFound(t *testing.T) {
	repo := newTestRepo()
	ctx := context.Background()

	_, err := repo.Get(ctx, "non-existent-id")

	assert.ErrorIs(t, err, ErrUserNotFound)
}

func TestUserRepository_Update(t *testing.T) {
	repo := newTestRepo()
	ctx := context.Background()

	userID, err := repo.Create(ctx, models.User{
		Name:     "old",
		Email:    "old@example.com",
		Password: "secret123",
	})
	require.NoError(t, err)

	err = repo.Update(ctx, models.User{
		ID:       userID,
		Name:     "new",
		Email:    "new@example.com",
		Password: "secret123",
	})
	require.NoError(t, err)

	user, err := repo.Get(ctx, userID)
	require.NoError(t, err)
	assert.Equal(t, "new", user.Name)
	assert.Equal(t, "new@example.com", user.Email)
}

func TestUserRepository_UpdateNotFound(t *testing.T) {
	repo := newTestRepo()
	ctx := context.Background()

	err := repo.Update(ctx, models.User{
		ID:       "non-existent-id",
		Name:     "new",
		Email:    "new@example.com",
		Password: "secret123",
	})

	assert.ErrorIs(t, err, ErrUserNotFound)
}

func TestUserRepository_Delete(t *testing.T) {
	repo := newTestRepo()
	ctx := context.Background()

	userID, err := repo.Create(ctx, models.User{
		Name:     "test",
		Email:    "test@example.com",
		Password: "secret123",
	})
	require.NoError(t, err)

	err = repo.Delete(ctx, userID)
	require.NoError(t, err)

	_, err = repo.Get(ctx, userID)
	assert.ErrorIs(t, err, ErrUserNotFound)
}

func TestUserRepository_DeleteNotFound(t *testing.T) {
	repo := newTestRepo()
	ctx := context.Background()

	err := repo.Delete(ctx, "non-existent-id")

	assert.ErrorIs(t, err, ErrUserNotFound)
}
