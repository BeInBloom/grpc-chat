package handler

import (
	"context"

	authv1 "github.com/BeInBloom/grpc-chat/gen/go/auth/v1"
	"github.com/BeInBloom/grpc-chat/services/auth/internal/models"
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
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		return nil, err
	}

	return &authv1.CreateResponse{Id: id}, nil
}

func (h *UserHandler) Get(ctx context.Context, req *authv1.GetRequest) (*authv1.GetResponse, error) {
	user, err := h.service.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &authv1.GetResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
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
		return nil, err
	}

	return &authv1.UpdateResponse{}, nil
}

func (h *UserHandler) Delete(ctx context.Context, req *authv1.DeleteRequest) (*authv1.DeleteResponse, error) {
	if err := h.service.Delete(ctx, req.GetId()); err != nil {
		return nil, err
	}

	return &authv1.DeleteResponse{}, nil
}
