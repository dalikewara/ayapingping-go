package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Connect connects to MySQL database.
// `dialect` usually consisting of at least a databases name and connection information.
func Connect(dialect string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dialect)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return db, err
	}
	return db, nil
}
