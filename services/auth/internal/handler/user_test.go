package handler

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	authv1 "github.com/BeInBloom/grpc-chat/gen/go/auth/v1"
	"github.com/BeInBloom/grpc-chat/services/auth/internal/handler/mocks"
	"github.com/BeInBloom/grpc-chat/services/auth/internal/models"
	"github.com/BeInBloom/grpc-chat/services/auth/internal/repository"
)

var testUUID = uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")

func TestUserHandler_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockuserService(ctrl)
	handler := New(mockService)

	ctx := context.Background()
	req := &authv1.CreateRequest{
		Name:     "test",
		Email:    "test@example.com",
		Password: "secret123",
		Role:     authv1.UserRole_USER_ROLE_USER,
	}

	mockService.EXPECT().
		Create(ctx, models.User{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
			Role:     int32(req.Role),
		}).
		Return(testUUID, nil)

	resp, err := handler.Create(ctx, req)

	require.NoError(t, err)
	assert.Equal(t, testUUID.String(), resp.GetId())
}

func TestUserHandler_CreateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockuserService(ctrl)
	handler := New(mockService)

	ctx := context.Background()
	req := &authv1.CreateRequest{
		Name:     "test",
		Email:    "test@example.com",
		Password: "secret123",
	}

	mockService.EXPECT().
		Create(ctx, gomock.Any()).
		Return(uuid.Nil, assert.AnError)

	resp, err := handler.Create(ctx, req)

	assert.Nil(t, resp)
	assert.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.Internal, st.Code())
}

func TestUserHandler_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockuserService(ctrl)
	handler := New(mockService)

	ctx := context.Background()
	req := &authv1.GetRequest{Id: testUUID.String()}

	mockService.EXPECT().
		Get(ctx, testUUID).
		Return(models.User{
			ID:    testUUID,
			Name:  "test",
			Email: "test@example.com",
			Role:  int32(authv1.UserRole_USER_ROLE_ADMIN),
		}, nil)

	resp, err := handler.Get(ctx, req)

	require.NoError(t, err)
	assert.Equal(t, testUUID.String(), resp.GetId())
	assert.Equal(t, "test", resp.GetName())
	assert.Equal(t, "test@example.com", resp.GetEmail())
	assert.Equal(t, authv1.UserRole_USER_ROLE_ADMIN, resp.GetRole())
}

func TestUserHandler_GetInvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockuserService(ctrl)
	handler := New(mockService)

	ctx := context.Background()
	req := &authv1.GetRequest{Id: "not-a-uuid"}

	resp, err := handler.Get(ctx, req)

	assert.Nil(t, resp)
	st, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, st.Code())
}

func TestUserHandler_GetNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockuserService(ctrl)
	handler := New(mockService)

	ctx := context.Background()
	nonExistent := uuid.MustParse("00000000-0000-0000-0000-000000000001")
	req := &authv1.GetRequest{Id: nonExistent.String()}

	mockService.EXPECT().
		Get(ctx, nonExistent).
		Return(models.User{}, repository.ErrUserNotFound)

	resp, err := handler.Get(ctx, req)

	assert.Nil(t, resp)
	assert.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.NotFound, st.Code())
}

func TestUserHandler_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockuserService(ctrl)
	handler := New(mockService)

	ctx := context.Background()
	req := &authv1.UpdateRequest{
		Id:    testUUID.String(),
		Name:  proto.String("updated"),
		Email: proto.String("updated@example.com"),
	}

	mockService.EXPECT().
		Update(ctx, models.User{
			ID:    testUUID,
			Name:  "updated",
			Email: "updated@example.com",
		}).
		Return(nil)

	resp, err := handler.Update(ctx, req)

	require.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestUserHandler_UpdateNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockuserService(ctrl)
	handler := New(mockService)

	ctx := context.Background()
	nonExistent := uuid.MustParse("00000000-0000-0000-0000-000000000001")
	req := &authv1.UpdateRequest{
		Id:   nonExistent.String(),
		Name: proto.String("updated"),
	}

	mockService.EXPECT().
		Update(ctx, gomock.Any()).
		Return(repository.ErrUserNotFound)

	resp, err := handler.Update(ctx, req)

	assert.Nil(t, resp)
	st, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.NotFound, st.Code())
}

func TestUserHandler_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockuserService(ctrl)
	handler := New(mockService)

	ctx := context.Background()
	req := &authv1.DeleteRequest{Id: testUUID.String()}

	mockService.EXPECT().
		Delete(ctx, testUUID).
		Return(nil)

	resp, err := handler.Delete(ctx, req)

	require.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestUserHandler_DeleteNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockuserService(ctrl)
	handler := New(mockService)

	ctx := context.Background()
	nonExistent := uuid.MustParse("00000000-0000-0000-0000-000000000001")
	req := &authv1.DeleteRequest{Id: nonExistent.String()}

	mockService.EXPECT().
		Delete(ctx, nonExistent).
		Return(repository.ErrUserNotFound)

	resp, err := handler.Delete(ctx, req)

	assert.Nil(t, resp)
	st, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.NotFound, st.Code())
}
