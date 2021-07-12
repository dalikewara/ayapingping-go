package env

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config object.
type Config struct {
	Key string `envconfig:"KEY"`
}

// NewEnv generates new dotenv configurations.
func NewEnv() (Config, error) {
	_ = godotenv.Load()
	cfg := Config{}
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
