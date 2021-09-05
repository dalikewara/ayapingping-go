package env

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Variable struct.
type Variable struct {
	Key string `envconfig:"KEY"`
}

// Generate generates new dotenv environment variables based on Variable struct.
func Generate() (Variable, error) {
	_ = godotenv.Load()
	cfg := Variable{}
	if err := envconfig.Process("", &cfg); err != nil {
		return Variable{}, err
	}
	return cfg, nil
}
