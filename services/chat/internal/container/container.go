package container

import (
	"log/slog"

	"github.com/BeInBloom/grpc-chat/pkg/logger"
	"github.com/BeInBloom/grpc-chat/services/chat/internal/app"
	"github.com/BeInBloom/grpc-chat/services/chat/internal/config"
)

type container struct {
	app    *app.App
	logger *slog.Logger
	config config.Config
}

func New(cfg config.Config) *container {
	return &container{
		config: cfg,
	}
}

func (c *container) App() *app.App {
	if c.app == nil {
		c.app = app.New()
	}

	return c.app
}

func (c *container) Logger() *slog.Logger {
	if c.logger == nil {
		c.logger = logger.New(c.config.Logger)
	}

	return c.logger
}
