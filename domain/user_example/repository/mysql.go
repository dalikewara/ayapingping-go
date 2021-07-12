package repository

import (
	"database/sql"
	"github.com/dalikewara/ayapingping-go/domain/user_example"
	"time"
)

// MySQL object.
type MySQL struct {
	db  *sql.DB
}

// NewMySQL generates new `MySQL` repository that implements `user_example.Repository`.
func NewMySQL(db *sql.DB) user_example.Repository {
	return &MySQL{
		db:  db,
	}
}

// GetByUsernamePassword gets user_example data by username and password.
func (m *MySQL) GetByUsernamePassword(request user_example.RepositoryGetByUsernamePasswordRequest) user_example.RepositoryGetByUsernamePasswordResponse {
	// Communicate to the real database. Example:
	// usr := &user_example.Entity{}
	// err := m.db.QueryRow("SELECT id, username, password FROM users WHERE username = ? AND password = ?",
	//	 request.Username, request.Password).Scan(&usr.Id, &usr.Username, &usr.Password)
	return user_example.RepositoryGetByUsernamePasswordResponse{
		User:  &user_example.Entity{
			Username: request.Username,
			Password: request.Password,
		},
		Error: nil,
	}
}

// GetByUsernamePasswordContext gets user_example data with context by username and password.
func (m *MySQL) GetByUsernamePasswordContext(request user_example.RepositoryGetByUsernamePasswordContextRequest) user_example.RepositoryGetByUsernamePasswordContextResponse {
	// Communicate to the real database. Example:
	// usr := &user_example.Entity{}
	// err := m.db.QueryRowContext(request.Ctx, "SELECT id, username, password FROM users WHERE username = ? AND password = ?",
	//	 request.Username, request.Password).Scan(&usr.Id, &usr.Username, &usr.Password)
	time.Sleep(2*time.Second)
	select {
		case <-request.Ctx.Done():
			return user_example.RepositoryGetByUsernamePasswordContextResponse{
				User: nil,
				Error: request.Ctx.Err(),
			}
	}
}
