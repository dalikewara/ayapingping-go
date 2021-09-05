package example

import (
	"database/sql"
)

// NewMySQLRepositoryParam parameter.
type NewMySQLRepositoryParam struct {
	Db *sql.DB
}

// RepositoryFindByIdParam parameter.
type RepositoryFindByIdParam struct {
	Id int64
	// Ctx context.Context  // If you want to use context.
}

// RepositoryFindByIdResult result.
type RepositoryFindByIdResult struct {
	Example *Example
	Error error
}

// RepositoryUpdateNameByIdParam parameter.
type RepositoryUpdateNameByIdParam struct {
	Id int64
	Name string
	// Ctx context.Context  // If you want to use context.
}

// RepositoryUpdateNameByIdResult result.
type RepositoryUpdateNameByIdResult struct {
	Error error
}
