package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	authv1 "github.com/BeInBloom/grpc-chat/gen/go/auth/v1"
	"github.com/BeInBloom/grpc-chat/services/auth/internal/handler"
)

const grpcPort = ":50051"

func main() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	authv1.RegisterUserAPIServiceServer(grpcServer, handler.NewUserService())

	reflection.Register(grpcServer)

	log.Printf("auth service listening on %s", grpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
