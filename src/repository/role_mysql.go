package repository

import (
	"database/sql"
	"github.com/dalikewara/ayapingping-go/src/entity"
	"time"
)

type roleMySQL struct {
	db *sql.DB
}

// NewRoleMySQL generates new roleMySQL that implements Role.
func NewRoleMySQL(param NewRoleMySQLParam) Role {
	return &roleMySQL{
		db: param.DB,
	}
}

// FindByUserID finds role data by user id from MySQL database.
func (r *roleMySQL) FindByUserID(param RoleFindByUserIDParam) RoleFindByUserIDResult {
	var result RoleFindByUserIDResult

	if param.UserId != 1 {
		result.Error = ErrorRoleNoDataFound
		return result
	}

	result.Role = &entity.Role{
		Id:        1,
		UserId:    param.UserId,
		Role:      "admin",
		CreatedAt: time.Now(),
	}

	return result
}
