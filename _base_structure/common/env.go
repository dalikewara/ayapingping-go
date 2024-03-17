package common

import (
	"errors"
	"fmt"
	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

type Env struct {
	AppEnv      string `env:"APP_ENV" envDefault:"development"`
	RESTPort    string `env:"REST_PORT" envDefault:"8080"`
	MySQLHost   string `env:"MYSQL_HOST" envDefault:"localhost"`
	MySQLPort   string `env:"MYSQL_PORT" envDefault:"3306"`
	MySQLUser   string `env:"MYSQL_USER" envDefault:""`
	MySQLPass   string `env:"MYSQL_PASS" envDefault:""`
	MySQLDBName string `env:"MYSQL_DB_NAME" envDefault:""`
}

func ParseEnv() (*Env, error) {
	envCfg := &Env{}

	if err := env.Parse(envCfg); err != nil {
		return nil, errors.New(fmt.Sprintf("error parsing env variables: %v", err.Error()))
	}

	return envCfg, nil
}
