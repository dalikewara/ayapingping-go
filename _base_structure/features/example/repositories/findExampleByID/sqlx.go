package findExampleByID

import (
	"context"
	"github.com/dalikewara/ayapingping-go/v4/_base_structure/domain"
	"github.com/jmoiron/sqlx"
)

type mysql struct {
	db *sqlx.DB
}

func NewMySQL(db *sqlx.DB) domain.FindExampleByIDRepository {
	return &mysql{
		db: db,
	}
}

func (m *mysql) ExecCtx(ctx context.Context, id uint64) (*domain.Example, error) {
	example := &domain.Example{
		ID:       id,
		Username: "dalikewara",
		Password: "password",
	}

	example.SetCreatedAtNow()

	return example, nil
}
