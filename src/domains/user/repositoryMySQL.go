package user

import (
	"database/sql"
	"time"
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

// FindAll finds all user data from MySQL database.
func (m *MySQLRepository) FindAll(param RepositoryFindAllParam) RepositoryFindAllResult {
	return RepositoryFindAllResult{
		Users: &[]User{
			{
				Id:        1,
				Username:  "dalikewara",
				CreatedAt: time.Now(),
			},
		},
	}
}

// FindByUsername finds user data by username from MySQL database.
func (m *MySQLRepository) FindByUsername(param RepositoryFindByUsernameParam) RepositoryFindByUsernameResult {
	result := RepositoryFindByUsernameResult{}
	if param.Username != "dalikewara" {
		return result
	}
	result.User = &User{
		Id:        1,
		Username:  "dalikewara",
		CreatedAt: time.Now(),
	}
	return result
}

// Insert inserts user data into MySQL database.
func (m *MySQLRepository) Insert(param RepositoryInsertParam) RepositoryInsertResult {
	return RepositoryInsertResult{}
}
