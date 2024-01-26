package findExampleByID

import (
	"context"
	"database/sql"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/entity/example"
)

type mysqlDB struct {
	db *sql.DB
}

// ExecCtx executes the repository main logic with context
func (s *mysqlDB) ExecCtx(ctx context.Context, id uint64) (*example.Model, error) {
	// Communicates with the real MySQL connection here, like querying.

	exampleModel := &example.Model{
		ID:   1,
		Name: "Dali Kewara",
	}

	exampleModel.SetCreatedAtNowUTC()

	return exampleModel, nil
}

// NewMySQLDB creates new MySQL database repository
func NewMySQLDB(db *sql.DB) Contract {
	return &mysqlDB{
		db: db,
	}
}
