package handler

import (
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	authv1 "github.com/BeInBloom/grpc-chat/gen/go/auth/v1"
	"github.com/BeInBloom/grpc-chat/services/auth/internal/models"
)

func toUserID(id string) (uuid.UUID, error) {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, status.Errorf(codes.InvalidArgument, "invalid user id: %s", err)
	}

	return parsed, nil
}

func toUser(req *authv1.CreateRequest) models.User {
	return models.User{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Role:     int32(req.GetRole()),
	}
}

func toUserUpdate(req *authv1.UpdateRequest, id uuid.UUID) models.User {
	user := models.User{ID: id}

	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		user.Email = *req.Email
	}

	return user
}

func toProtoGetResponse(user models.User) *authv1.GetResponse {
	return &authv1.GetResponse{
		Id:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		Role:      authv1.UserRole(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
