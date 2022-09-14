package env

import (
	"github.com/joho/godotenv"
	"os"
)

var _ = godotenv.Load()

var Example = os.Getenv("EXAMPLE")
