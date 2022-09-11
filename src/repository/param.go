package repository

import (
	"context"
	"database/sql"
	"github.com/dalikewara/ayapingping-go/v2/src/entity"
	"github.com/dalikewara/ayapingping-go/v2/src/library/errs"
)

type NewParam struct {
	MySQLDB *sql.DB
}

type NewUserMySQLParam struct {
	DB *sql.DB
}

type UserFindAllParam struct {
	Ctx context.Context // In case you need a context.
}

type UserFindAllResult struct {
	Users []*entity.User `json:"users"`
	Error errs.Errs      `json:"error"`
}

type NewRoleMySQLParam struct {
	DB *sql.DB
}

type RoleFindByUserIDParam struct {
	Ctx    context.Context
	UserId int `json:"user_id"`
}

type RoleFindByUserIDResult struct {
	Role  *entity.Role `json:"role"`
	Error errs.Errs    `json:"error"`
}
