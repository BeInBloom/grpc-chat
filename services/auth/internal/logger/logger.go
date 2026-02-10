package logger

import (
	"log/slog"
	"os"
	"strings"

	"github.com/BeInBloom/grpc-chat/services/auth/internal/models"
)

func New(config models.LoggerConfig) *slog.Logger {
	var handler slog.Handler

	var level slog.Level
	switch strings.ToLower(config.Logger.Level) {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: config.Logger.AddSource,
	}

	if config.Env == "local" {
		handler = slog.NewTextHandler(os.Stdout, opts)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	logger := slog.New(handler).With(
		slog.String("env", config.Env),
		slog.String("service", config.Service),
	)

	return logger
}
