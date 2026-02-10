package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/BeInBloom/grpc-chat/services/auth/internal/models"
	"github.com/BeInBloom/grpc-chat/services/auth/internal/services/mocks"
)

func TestUserService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockuserRepository(ctrl)
	service := New(mockRepo)

	ctx := context.Background()
	user := models.User{
		Name:     "test",
		Email:    "test@example.com",
		Password: "secret123",
	}

	mockRepo.EXPECT().
		Create(ctx, user).
		Return("uuid-123", nil)

	id, err := service.Create(ctx, user)

	require.NoError(t, err)
	assert.Equal(t, "uuid-123", id)
}

func TestUserService_CreateValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockuserRepository(ctrl)
	service := New(mockRepo)

	ctx := context.Background()

	_, err := service.Create(ctx, models.User{
		Name:     "",
		Email:    "test@example.com",
		Password: "secret123",
	})
	assert.Error(t, err)

	_, err = service.Create(ctx, models.User{
		Name:     "test",
		Email:    "",
		Password: "secret123",
	})
	assert.Error(t, err)

	_, err = service.Create(ctx, models.User{
		Name:     "test",
		Email:    "test@example.com",
		Password: "",
	})
	assert.Error(t, err)
}

func TestUserService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockuserRepository(ctrl)
	service := New(mockRepo)

	ctx := context.Background()
	expectedUser := models.User{
		ID:       "uuid-123",
		Name:     "test",
		Email:    "test@example.com",
		Password: "secret123",
	}

	mockRepo.EXPECT().
		Get(ctx, "uuid-123").
		Return(expectedUser, nil)

	user, err := service.Get(ctx, "uuid-123")

	require.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestUserService_GetNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockuserRepository(ctrl)
	service := New(mockRepo)

	ctx := context.Background()

	mockRepo.EXPECT().
		Get(ctx, "non-existent").
		Return(models.User{}, assert.AnError)

	_, err := service.Get(ctx, "non-existent")

	assert.Error(t, err)
}

func TestUserService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockuserRepository(ctrl)
	service := New(mockRepo)

	ctx := context.Background()

	existingUser := models.User{
		ID:       "uuid-123",
		Name:     "old",
		Email:    "old@example.com",
		Password: "secret123",
	}

	mockRepo.EXPECT().
		Get(ctx, "uuid-123").
		Return(existingUser, nil)

	mockRepo.EXPECT().
		Update(ctx, models.User{
			ID:       "uuid-123",
			Name:     "updated",
			Email:    "updated@example.com",
			Password: "secret123",
		}).
		Return(nil)

	err := service.Update(ctx, models.User{
		ID:    "uuid-123",
		Name:  "updated",
		Email: "updated@example.com",
	})

	require.NoError(t, err)
}

func TestUserService_UpdatePartial(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockuserRepository(ctrl)
	service := New(mockRepo)

	ctx := context.Background()

	existingUser := models.User{
		ID:       "uuid-123",
		Name:     "old",
		Email:    "old@example.com",
		Password: "secret123",
	}

	mockRepo.EXPECT().
		Get(ctx, "uuid-123").
		Return(existingUser, nil)

	mockRepo.EXPECT().
		Update(ctx, models.User{
			ID:       "uuid-123",
			Name:     "updated",
			Email:    "old@example.com",
			Password: "secret123",
		}).
		Return(nil)

	err := service.Update(ctx, models.User{
		ID:   "uuid-123",
		Name: "updated",
	})

	require.NoError(t, err)
}

func TestUserService_UpdateNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockuserRepository(ctrl)
	service := New(mockRepo)

	ctx := context.Background()

	mockRepo.EXPECT().
		Get(ctx, "non-existent").
		Return(models.User{}, assert.AnError)

	err := service.Update(ctx, models.User{
		ID:    "non-existent",
		Name:  "updated",
		Email: "updated@example.com",
	})

	assert.Error(t, err)
}

func TestUserService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockuserRepository(ctrl)
	service := New(mockRepo)

	ctx := context.Background()

	mockRepo.EXPECT().
		Delete(ctx, "uuid-123").
		Return(nil)

	err := service.Delete(ctx, "uuid-123")

	require.NoError(t, err)
}

func TestUserService_DeleteNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockuserRepository(ctrl)
	service := New(mockRepo)

	ctx := context.Background()

	mockRepo.EXPECT().
		Delete(ctx, "non-existent").
		Return(assert.AnError)

	err := service.Delete(ctx, "non-existent")

	assert.Error(t, err)
}
