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
}

// Connect connects to MySQL database.
func Connect(param ConnectParam) (*sql.DB, error) {
	dialect := param.User
	if param.Pass != "" {
		dialect = dialect + ":" + param.Pass
	}
	dialect = dialect + "@tcp(" + param.Host
	if param.Port != "" {
		dialect = dialect + ":" + param.Port
	}
	dialect = dialect + ")"
	if param.DBName != "" {
		dialect = dialect + "/" + param.DBName
	}
	db, err := sql.Open("mysql", dialect)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return db, err
	}
	return db, nil
}
