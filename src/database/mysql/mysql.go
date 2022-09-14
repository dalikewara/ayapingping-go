package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type ConnectParam struct {
	Host   string
	Port   string
	User   string
	Pass   string
	DBName string
	Option string
}

// Connect connects to MySQL database.
func Connect(param ConnectParam) (*sql.DB, error) {
	panic("implement me")
}
