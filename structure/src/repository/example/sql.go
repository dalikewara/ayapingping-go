package example

import (
	"context"
	"database/sql"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/entity/example"
)

type sqlDB struct {
	db *sql.DB
}

func (s *sqlDB) ExecuteCtx(ctx context.Context, id string) (*example.BaseModel, error) {
	panic("implement me")
}

// New

func NewSQLite(db *sql.DB) Contract {
	return &sqlDB{
		db: db,
	}
}
