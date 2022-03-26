// Example database package.

package mysql

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
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
	dialect := param.User
	if param.Pass != "" {
		dialect += ":" + param.Pass
	}
	dialect += "@tcp(" + param.Host
	if param.Port != "" {
		dialect += ":" + param.Port
	}
	dialect += ")"
	if param.DBName != "" {
		dialect += "/" + param.DBName
	}
	dialect += param.Option
	db, err := sql.Open("mysql", dialect)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return db, err
	}
	return db, nil
}

// ConnectMock mocks the MySQL database connection.
func ConnectMock() (*sql.DB, sqlmock.Sqlmock, error) {
	db, sqlMock, err := sqlmock.New()
	return db, sqlMock, err
}
