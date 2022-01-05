package example

import (
	"context"
	"database/sql"
)

// NewMySQLRepositoryParam is the argument sets of the NewMySQLRepository function.
type NewMySQLRepositoryParam struct {
	Db *sql.DB
}

// RepositoryFindByIdParam is the argument sets of the RepositoryInterface.FindById method.
type RepositoryFindByIdParam struct {
	Id  int64
	Ctx *context.Context // If you want to use context.
}

// RepositoryFindByIdResult is the response model of the RepositoryInterface.FindById method.
type RepositoryFindByIdResult struct {
	Example *Example
	Error   error
}

// RepositoryUpdateNameByIdParam is the argument sets of the RepositoryInterface.UpdateNameById method.
type RepositoryUpdateNameByIdParam struct {
	Id   int64
	Name string
	Ctx  *context.Context // If you want to use context.
}

// RepositoryUpdateNameByIdResult is the response model of the RepositoryInterface.UpdateNameById method.
type RepositoryUpdateNameByIdResult struct {
	Error error
}
