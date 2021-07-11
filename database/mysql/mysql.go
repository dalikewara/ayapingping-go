package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// NewMySQL opens MySQL database connection.
// `dialect` usually consisting of at least a databases name and connection information.
func NewMySQL(dialect string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dialect)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return db, err
	}
	return db, nil
}
