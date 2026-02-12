package main

import (
	"context"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/BeInBloom/grpc-chat/services/auth/internal/config"
	"github.com/BeInBloom/grpc-chat/services/auth/internal/container"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	config := config.New()
	c := container.New(config)
	log := c.Logger()

	log.Info("starting auth app...")

	a := c.App()
	if err := a.Run(ctx); err != nil {
		log.Error("runtime error", slog.String("error", err.Error()))
	}
}
