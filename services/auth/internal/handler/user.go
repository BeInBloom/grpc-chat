package handler

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	authv1 "github.com/BeInBloom/grpc-chat/gen/go/auth/v1"
	"github.com/BeInBloom/grpc-chat/services/auth/internal/models"
	"github.com/BeInBloom/grpc-chat/services/auth/internal/repository"
)

//go:generate mockgen -source=user.go -destination=mocks/mock_service.go -package=mocks

type userService interface {
	Create(ctx context.Context, user models.User) (uuid.UUID, error)
	Get(ctx context.Context, id uuid.UUID) (models.User, error)
	Update(ctx context.Context, user models.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type UserHandler struct {
	authv1.UnimplementedUserAPIServiceServer
	service userService
}

func New(service userService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Create(ctx context.Context, req *authv1.CreateRequest) (*authv1.CreateResponse, error) {
	id, err := h.service.Create(ctx, toUser(req))
	if err != nil {
		return nil, toGRPCError(err)
	}

	return &authv1.CreateResponse{Id: id.String()}, nil
}

func (h *UserHandler) Get(ctx context.Context, req *authv1.GetRequest) (*authv1.GetResponse, error) {
	id, err := toUserID(req.GetId())
	if err != nil {
		return nil, err
	}

	user, err := h.service.Get(ctx, id)
	if err != nil {
		return nil, toGRPCError(err)
	}

	return toProtoGetResponse(user), nil
}

func (h *UserHandler) Update(ctx context.Context, req *authv1.UpdateRequest) (*authv1.UpdateResponse, error) {
	id, err := toUserID(req.GetId())
	if err != nil {
		return nil, err
	}

	if err := h.service.Update(ctx, toUserUpdate(req, id)); err != nil {
		return nil, toGRPCError(err)
	}

	return &authv1.UpdateResponse{}, nil
}

func (h *UserHandler) Delete(ctx context.Context, req *authv1.DeleteRequest) (*authv1.DeleteResponse, error) {
	id, err := toUserID(req.GetId())
	if err != nil {
		return nil, err
	}

	if err := h.service.Delete(ctx, id); err != nil {
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
