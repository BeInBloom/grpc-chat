package app

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	authv1 "github.com/BeInBloom/grpc-chat/gen/go/auth/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	handlers authv1.UserAPIServiceServer
	logger   *slog.Logger
	addr     string
}

func New(
	addr string,
	logger *slog.Logger,
	handlers authv1.UserAPIServiceServer,
) *App {
	logger = logger.With("layer", "auth app")

	return &App{
		logger:   logger,
		handlers: handlers,
		addr:     addr,
	}
}

func (a *App) Run(ctx context.Context) error {
	lis, err := net.Listen("tcp", a.addr)
	if err != nil {
		return fmt.Errorf("running fail: %w", err)
	}

	grpcServer := grpc.NewServer()
	authv1.RegisterUserAPIServiceServer(grpcServer, a.handlers)
	reflection.Register(grpcServer)

	a.logger.Info("auth service listening", slog.String("addr", a.addr))

	errCh := make(chan error, 1)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			errCh <- fmt.Errorf("something wrong: %w", err)
		}
	}()

	select {
	case <-ctx.Done():
		a.logger.Info("auth service shutdown by context")
		grpcServer.GracefulStop()
		a.logger.Info("auth service stopped")
		return nil
	case err := <-errCh:
		return fmt.Errorf("auth service failed: %w", err)
	}
}
