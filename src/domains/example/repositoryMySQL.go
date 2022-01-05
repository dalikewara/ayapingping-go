package example

import (
	"database/sql"
)

type MySQLRepository struct {
	db *sql.DB
}

// NewMySQLRepository generates new MySQLRepository.
func NewMySQLRepository(param NewMySQLRepositoryParam) RepositoryInterface {
	return &MySQLRepository{
		db: param.Db,
	}
}

// FindById finds example data from database by id.
func (m *MySQLRepository) FindById(param RepositoryFindByIdParam) RepositoryFindByIdResult {
	// Repository is a place where you communicate with the real database, example:
	// ex := &Example{}
	// err := m.db.QueryRow("SELECT id FROM examples").Scan(&ex.Id)
	return RepositoryFindByIdResult{
		Example: &Example{
			Id:   param.Id,
			Name: "John Doe",
		},
		Error: nil,
	}
}

// UpdateNameById updates example name on database by id.
func (m *MySQLRepository) UpdateNameById(param RepositoryUpdateNameByIdParam) RepositoryUpdateNameByIdResult {
	// Repository is a place where you communicate with the real database.
	return RepositoryUpdateNameByIdResult{
		Error: nil,
	}
}
