package handler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/proto"

	authv1 "github.com/BeInBloom/grpc-chat/gen/go/auth/v1"
	"github.com/BeInBloom/grpc-chat/services/auth/internal/handler/mocks"
	"github.com/BeInBloom/grpc-chat/services/auth/internal/models"
)

func TestUserHandler_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockuserService(ctrl)
	handler := New(mockService)

	ctx := context.Background()
	req := &authv1.CreateRequest{
		Name:  "test",
		Email: "test@example.com",
	}

	mockService.EXPECT().
		Create(ctx, models.User{
			Name:  req.Name,
			Email: req.Email,
		}).
		Return("uuid-123", nil)

	resp, err := handler.Create(ctx, req)

	require.NoError(t, err)
	assert.Equal(t, "uuid-123", resp.GetId())
}

func TestUserHandler_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockuserService(ctrl)
	handler := New(mockService)

	ctx := context.Background()
	req := &authv1.GetRequest{Id: "uuid-123"}

	mockService.EXPECT().
		Get(ctx, "uuid-123").
		Return(models.User{
			ID:    "uuid-123",
			Name:  "test",
			Email: "test@example.com",
		}, nil)

	resp, err := handler.Get(ctx, req)

	require.NoError(t, err)
	assert.Equal(t, "uuid-123", resp.GetId())
	assert.Equal(t, "test", resp.GetName())
	assert.Equal(t, "test@example.com", resp.GetEmail())
}

func TestUserHandler_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockuserService(ctrl)
	handler := New(mockService)

	ctx := context.Background()
	req := &authv1.UpdateRequest{
		Id:    "uuid-123",
		Name:  proto.String("updated"),
		Email: proto.String("updated@example.com"),
	}

	mockService.EXPECT().
		Update(ctx, models.User{
			ID:    "uuid-123",
			Name:  "updated",
			Email: "updated@example.com",
		}).
		Return(nil)

	resp, err := handler.Update(ctx, req)

	require.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestUserHandler_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockuserService(ctrl)
	handler := New(mockService)

	ctx := context.Background()
	req := &authv1.DeleteRequest{Id: "uuid-123"}

	mockService.EXPECT().
		Delete(ctx, "uuid-123").
		Return(nil)

	resp, err := handler.Delete(ctx, req)

	require.NoError(t, err)
	assert.NotNil(t, resp)
}
