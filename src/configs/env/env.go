package env

import (
	"github.com/joho/godotenv"
	"os"
)

var _ = godotenv.Load()

var AppEnv = os.Getenv("APP_ENV")
