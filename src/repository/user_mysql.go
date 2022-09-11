package repository

import (
	"database/sql"
	"github.com/dalikewara/ayapingping-go/v2/src/entity"
	"time"
)

type userMySQL struct {
	db *sql.DB
}

// NewUserMySQL generates new userMySQL that implements User.
func NewUserMySQL(param NewUserMySQLParam) User {
	return &userMySQL{
		db: param.DB,
	}
}

// FindAll finds all user data from MySQL database.
func (r *userMySQL) FindAll(param UserFindAllParam) UserFindAllResult {
	var result UserFindAllResult
	var users []*entity.User

	users = append(users, &entity.User{
		Id:        1,
		Username:  "johndoe",
		CreatedAt: time.Now(),
	})
	result.Users = users

	return result
}
