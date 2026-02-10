package container

import (
	"log/slog"

	"github.com/BeInBloom/grpc-chat/services/chat/internal/app"
	"github.com/BeInBloom/grpc-chat/services/chat/internal/config"
	"github.com/BeInBloom/grpc-chat/services/chat/internal/logger"
)

type container struct {
	app    *app.App
	logger *slog.Logger
	config config.Config
}

func New(config config.Config) *container {
	return &container{
		config: config,
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
		c.logger = logger.New()
	}

	return c.logger
}
