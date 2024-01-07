package config

import (
	"github.com/dalikewara/ayapingping-go/v3/structure/src/config/constant"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/config/env"
)

type Config struct {
	RESTPort        string `env:"REST_PORT" envDefault:"8080"`
	MySQLConnString string `env:"MYSQL_CONN_STRING"`
}

// InitConfig initializes configurations
func InitConfig() (*Config, error) {
	cfg := &Config{}

	// Parse env variables
	if err := env.ParseEnv(cfg); err != nil {
		return nil, err
	}

	if cfg.RESTPort == "" {
		cfg.RESTPort = constant.RESTPort
	}

	return cfg, nil
}
