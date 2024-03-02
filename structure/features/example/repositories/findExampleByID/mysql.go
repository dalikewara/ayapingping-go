package findExampleByID

import (
	"context"
	"database/sql"
	"github.com/dalikewara/ayapingping-go/v4/structure/domain"
)

type mysql struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) domain.FindExampleByIDRepository {
	return &mysql{
		db: db,
	}
}

func (m *mysql) ExecCtx(ctx context.Context, id uint64) (*domain.Example, error) {
	example := &domain.Example{
		ID:       1,
		Username: "dalikewara",
		Password: "dalikewara123",
	}

	example.SetCreatedAtNow()

	return example, nil
}
