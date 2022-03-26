package user

import (
	"context"
	"database/sql"
)

type NewMySQLRepositoryParam struct {
	Db *sql.DB
}

type RepositoryFindAllParam struct {
	Ctx context.Context // If you want to use context.
}

type RepositoryFindAllResult struct {
	Users *[]User `json:"users"`
	Error error   `json:"error"`
}

type RepositoryFindByUsernameParam struct {
	Username string          `json:"username"`
	Ctx      context.Context // If you want to use context.
}

type RepositoryFindByUsernameResult struct {
	User  *User `json:"user"`
	Error error `json:"error"`
}

type RepositoryInsertParam struct {
	Username string          `json:"username"`
	Ctx      context.Context // If you want to use context.
}

type RepositoryInsertResult struct {
	Error error `json:"error"`
}
