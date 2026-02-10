package config

import (
	"log"
	"os"

	"github.com/BeInBloom/grpc-chat/services/auth/internal/models"
	"github.com/ilyakaznacheev/cleanenv"
)

func New() models.Config {
	configPath := ".env"

	if path := os.Getenv("CONFIG_PATH"); path != "" {
		configPath = path
	}

	var config models.Config

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
