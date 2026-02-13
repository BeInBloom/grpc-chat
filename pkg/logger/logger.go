package logger

import (
	"log/slog"
	"os"
	"strings"
)

type Config struct {
	Env     string `yaml:"env" env:"ENV" env-default:"local"`
	Service string `yaml:"service" env:"SERVICE_NAME" env-default:"my-app"`
	Logger  struct {
		Level     string `yaml:"level" env:"LOG_LEVEL" env-default:"info"`
		AddSource bool   `yaml:"add_source" env:"LOG_ADD_SOURCE" env-default:"false"`
	}
}

func New(config Config) *slog.Logger {
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
