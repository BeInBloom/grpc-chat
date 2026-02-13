package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/BeInBloom/grpc-chat/pkg/logger"
)

type Config struct {
	Addr   string        `yaml:"addr" env:"addr" env-default:"localhost:50051"`
	Logger logger.Config `yaml:"logger"`
}

func New() Config {
	configPath := ".env"

	if path := os.Getenv("CONFIG_PATH"); path != "" {
		configPath = path
	}

	var config Config

	if _, err := os.Stat(configPath); err == nil {
		if err := cleanenv.ReadConfig(configPath, &config); err != nil {
			log.Fatalf("cannot read config: %s", err)
		}
	} else {
		if err := cleanenv.ReadEnv(&config); err != nil {
			log.Fatalf("cannot read env: %s", err)
		}
	}

	return config
}
