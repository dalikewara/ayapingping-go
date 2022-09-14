package repository

import (
	"context"
	"database/sql"
	"github.com/dalikewara/ayapingping-go/v2/src/entity"
)

type NewParam struct {
	MySQLDB *sql.DB
}

type NewExampleMySQLParam struct {
	DB *sql.DB
}

type ExampleFindAllParam struct {
	Ctx context.Context // In case you need a context.
}

type ExampleFindAllResult struct {
	Examples []*entity.Example `json:"examples"`
	Error    error             `json:"error"`
}
