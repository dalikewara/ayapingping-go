package env

type Env struct {
	AppEnv   string `env:"APP_ENV" envDefault:"development"`
	RESTPort string `env:"REST_PORT" envDefault:"8080"`
}
