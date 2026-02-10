package handler

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	authv1 "github.com/BeInBloom/grpc-chat/gen/go/auth/v1"
	"github.com/BeInBloom/grpc-chat/services/auth/internal/models"
	"github.com/BeInBloom/grpc-chat/services/auth/internal/repository"
)

//go:generate mockgen -source=user.go -destination=mocks/mock_service.go -package=mocks

type userService interface {
	Create(ctx context.Context, user models.User) (string, error)
	Get(ctx context.Context, id string) (models.User, error)
	Update(ctx context.Context, user models.User) error
	Delete(ctx context.Context, id string) error
}

type UserHandler struct {
	authv1.UnimplementedUserAPIServiceServer
	service userService
}

func New(service userService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Create(ctx context.Context, req *authv1.CreateRequest) (*authv1.CreateResponse, error) {
	id, err := h.service.Create(ctx, models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     int32(req.Role),
	})
	if err != nil {
		return nil, toGRPCError(err)
	}

	return &authv1.CreateResponse{Id: id}, nil
}

func (h *UserHandler) Get(ctx context.Context, req *authv1.GetRequest) (*authv1.GetResponse, error) {
	user, err := h.service.Get(ctx, req.GetId())
	if err != nil {
		return nil, toGRPCError(err)
	}

	return &authv1.GetResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      authv1.UserRole(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}, nil
}

func (h *UserHandler) Update(ctx context.Context, req *authv1.UpdateRequest) (*authv1.UpdateResponse, error) {
	user := models.User{
		ID: req.GetId(),
	}
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		user.Email = *req.Email
	}

	if err := h.service.Update(ctx, user); err != nil {
		return nil, toGRPCError(err)
	}

	return &authv1.UpdateResponse{}, nil
}

func (h *UserHandler) Delete(ctx context.Context, req *authv1.DeleteRequest) (*authv1.DeleteResponse, error) {
	if err := h.service.Delete(ctx, req.GetId()); err != nil {
		return nil, toGRPCError(err)
	}

	return &authv1.DeleteResponse{}, nil
}

func toGRPCError(err error) error {
	if errors.Is(err, repository.ErrUserNotFound) {
		return status.Error(codes.NotFound, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}
