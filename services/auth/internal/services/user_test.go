package services

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/BeInBloom/grpc-chat/services/auth/internal/models"
	"github.com/BeInBloom/grpc-chat/services/auth/internal/services/mocks"
)

var testUUID = uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")

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
		Return(testUUID, nil)

	id, err := service.Create(ctx, user)

	require.NoError(t, err)
	assert.Equal(t, testUUID, id)
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
		ID:       testUUID,
		Name:     "test",
		Email:    "test@example.com",
		Password: "secret123",
	}

	mockRepo.EXPECT().
		Get(ctx, testUUID).
		Return(expectedUser, nil)

	user, err := service.Get(ctx, testUUID)

	require.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestUserService_GetNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockuserRepository(ctrl)
	service := New(mockRepo)

	ctx := context.Background()
	nonExistent := uuid.MustParse("00000000-0000-0000-0000-000000000001")

	mockRepo.EXPECT().
		Get(ctx, nonExistent).
		Return(models.User{}, assert.AnError)

	_, err := service.Get(ctx, nonExistent)

	assert.Error(t, err)
}

func TestUserService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockuserRepository(ctrl)
	service := New(mockRepo)

	ctx := context.Background()

	existingUser := models.User{
		ID:       testUUID,
		Name:     "old",
		Email:    "old@example.com",
		Password: "secret123",
	}

	mockRepo.EXPECT().
		Get(ctx, testUUID).
		Return(existingUser, nil)

	mockRepo.EXPECT().
		Update(ctx, models.User{
			ID:       testUUID,
			Name:     "updated",
			Email:    "updated@example.com",
			Password: "secret123",
		}).
		Return(nil)

	err := service.Update(ctx, models.User{
		ID:    testUUID,
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
		ID:       testUUID,
		Name:     "old",
		Email:    "old@example.com",
		Password: "secret123",
	}

	mockRepo.EXPECT().
		Get(ctx, testUUID).
		Return(existingUser, nil)

	mockRepo.EXPECT().
		Update(ctx, models.User{
			ID:       testUUID,
			Name:     "updated",
			Email:    "old@example.com",
			Password: "secret123",
		}).
		Return(nil)

	err := service.Update(ctx, models.User{
		ID:   testUUID,
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
	nonExistent := uuid.MustParse("00000000-0000-0000-0000-000000000001")

	mockRepo.EXPECT().
		Get(ctx, nonExistent).
		Return(models.User{}, assert.AnError)

	err := service.Update(ctx, models.User{
		ID:    nonExistent,
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
		Delete(ctx, testUUID).
		Return(nil)

	err := service.Delete(ctx, testUUID)

	require.NoError(t, err)
}

func TestUserService_DeleteNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockuserRepository(ctrl)
	service := New(mockRepo)

	ctx := context.Background()
	nonExistent := uuid.MustParse("00000000-0000-0000-0000-000000000001")

	mockRepo.EXPECT().
		Delete(ctx, nonExistent).
		Return(assert.AnError)

	err := service.Delete(ctx, nonExistent)

	assert.Error(t, err)
}
