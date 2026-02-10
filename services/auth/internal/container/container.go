package container

import (
	"log/slog"

	"github.com/BeInBloom/grpc-chat/services/auth/internal/app"
	"github.com/BeInBloom/grpc-chat/services/auth/internal/handler"
	"github.com/BeInBloom/grpc-chat/services/auth/internal/logger"
	"github.com/BeInBloom/grpc-chat/services/auth/internal/models"
	"github.com/BeInBloom/grpc-chat/services/auth/internal/repository"
	"github.com/BeInBloom/grpc-chat/services/auth/internal/services"
)

type container struct {
	config      models.Config
	userService *services.UserService
	userRepo    *repository.UserRepository
	logger      *slog.Logger
	handlers    *handler.UserHandler
	app         *app.App
}

func New(config models.Config) *container {
	return &container{config: config}
}

func (c *container) App() *app.App {
	if c.app == nil {
		c.app = app.New(c.config.Addr, c.Logger(), c.Handler())
	}

	return c.app
}

func (c *container) Handler() *handler.UserHandler {
	if c.handlers == nil {
		c.handlers = handler.New(c.UserService())
	}

	return c.handlers
}

func (c *container) UserService() *services.UserService {
	if c.userService == nil {
		c.userService = services.New(c.UserRepo())
	}

	return c.userService
}

func (c *container) Logger() *slog.Logger {
	if c.logger == nil {
		c.logger = logger.New(c.config.LoggerConfig)
	}

	return c.logger
}

func (c *container) UserRepo() *repository.UserRepository {
	if c.userRepo == nil {
		c.userRepo = repository.New(c.Config().UserRepository)
	}

	return c.userRepo
}

func (c *container) Config() models.Config {
	return c.config
}
