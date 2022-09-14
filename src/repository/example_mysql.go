package repository

import (
	"database/sql"
)

type exampleMySQL struct {
	db *sql.DB
}

// NewExampleMySQL generates new exampleMySQL that implements Example.
func NewExampleMySQL(param NewExampleMySQLParam) Example {
	panic("implement me")
}

// FindAll finds all example data from MySQL database.
func (r *exampleMySQL) FindAll(param ExampleFindAllParam) ExampleFindAllResult {
	panic("implement me")
}
