package example

import (
	"database/sql"
)

// MySQLRepository struct.
type MySQLRepository struct {
	db *sql.DB
}

// NewMySQLRepository generates new MySQLRepository.
func NewMySQLRepository(param NewMySQLRepositoryParam) RepositoryInterface {
	return &MySQLRepository{
		db: param.Db,
	}
}

// FindById finds `example` from database by id.
func (m *MySQLRepository) FindById(param RepositoryFindByIdParam) RepositoryFindByIdResult {
	// Communicate to the real database. Example:
	// ex := &Example{}
	// err := m.db.QueryRow("SELECT id FROM examples").Scan(&ex.Id)
	return RepositoryFindByIdResult{
		Example: &Example{
			Id: param.Id,
			Name: "John Doe",
		},
		Error: nil,
	}
}

// UpdateNameById updates `example` name to database by id.
func (m *MySQLRepository) UpdateNameById(param RepositoryUpdateNameByIdParam) RepositoryUpdateNameByIdResult {
	// Communicate to the real database.
	return RepositoryUpdateNameByIdResult{
		Error: nil,
	}
}
