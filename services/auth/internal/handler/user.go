package handler

import (
	"context"
	"log"

	authv1 "github.com/BeInBloom/grpc-chat/gen/go/auth/v1"
)

type UserService struct {
	authv1.UnimplementedUserAPIServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Create(ctx context.Context, req *authv1.CreateRequest) (*authv1.CreateResponse, error) {
	log.Printf("Create called: name=%s, email=%s", req.GetName(), req.GetEmail())
	return &authv1.CreateResponse{Id: 1}, nil
}

func (s *UserService) Get(ctx context.Context, req *authv1.GetRequest) (*authv1.GetResponse, error) {
	log.Printf("Get called: id=%d", req.GetId())
	return &authv1.GetResponse{
		Id:    req.GetId(),
		Name:  "stub_user",
		Email: "stub@example.com",
		Role:  authv1.UserRole_USER_ROLE_USER,
	}, nil
}

func (s *UserService) Update(ctx context.Context, req *authv1.UpdateRequest) (*authv1.UpdateResponse, error) {
	log.Printf("Update called: id=%d", req.GetId())
	return &authv1.UpdateResponse{}, nil
}

func (s *UserService) Delete(ctx context.Context, req *authv1.DeleteRequest) (*authv1.DeleteResponse, error) {
	log.Printf("Delete called: id=%d", req.GetId())
	return &authv1.DeleteResponse{}, nil
}
