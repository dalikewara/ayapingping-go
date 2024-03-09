package env

func Parse() (*Env, error) {
	return &Env{
		AppEnv:   "development",
		RESTPort: "8080",
	}, nil
}
