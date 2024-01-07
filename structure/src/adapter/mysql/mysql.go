package mysql

import "database/sql"

// NewMySQL creates new MySQL database connection
func NewMySQL(connString string) (*sql.DB, error) {
	// Connect to the real database here...
	return &sql.DB{}, nil
}
