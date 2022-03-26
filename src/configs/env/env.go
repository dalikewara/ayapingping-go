// Example config package.

package env

import (
	"github.com/joho/godotenv"
	"os"
)

var _ = godotenv.Load()

var AppEnv = os.Getenv("APP_ENV")
var AppPort = os.Getenv("APP_PORT")

var MySQLHost = os.Getenv("MYSQL_HOST")
var MySQLPort = os.Getenv("MYSQL_PORT")
var MySQLUser = os.Getenv("MYSQL_USER")
var MySQLPass = os.Getenv("MYSQL_PASS")
var MySQLDBName = os.Getenv("MYSQL_DBNAME")
