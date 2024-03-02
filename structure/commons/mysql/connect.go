package mysql

import (
	"database/sql"
)

func Connect() (*sql.DB, error) {
	return &sql.DB{}, nil
}
