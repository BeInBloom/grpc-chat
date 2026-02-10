package main

import (
	"context"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/BeInBloom/grpc-chat/services/chat/internal/config"
	"github.com/BeInBloom/grpc-chat/services/chat/internal/container"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := config.New()
	c := container.New(cfg)
	log := c.Logger()

	log.Info("starting chat app...")

	a := c.App()
	if err := a.Run(ctx); err != nil {
		log.Error("runtime error", slog.String("error", err.Error()))
	}
}
