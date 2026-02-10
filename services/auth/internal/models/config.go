package models

type (
	Config struct {
		UserRepository UserRepositoryConfig
		Addr           string `yaml:"addr" env:"addr" env-default:"localhost:50051"`
		LoggerConfig   LoggerConfig
	}

	UserRepositoryConfig struct{}

	LoggerConfig struct {
		Env     string `yaml:"env" env:"ENV" env-default:"local"`
		Service string `yaml:"service" env:"SERVICE_NAME" env-default:"my-app"`
		Logger  struct {
			Level     string `yaml:"level" env:"LOG_LEVEL" env-default:"info"`
			AddSource bool   `yaml:"add_source" env:"LOG_ADD_SOURCE" env-default:"false"`
		}
	}
)
