package service

import (
	"context"
	"github.com/dalikewara/ayapingping-go/v2/src/entity"
	"github.com/dalikewara/ayapingping-go/v2/src/library/errs"
	"github.com/dalikewara/ayapingping-go/v2/src/repository"
)

type Config struct {
	SystemUserRole string
	SystemUserIds  []int
}

type NewParam struct {
	Repo   *repository.Repository
	Config *Config
}

type NewUserV1Param struct {
	UserRepo repository.User
}

type UserGetAllParam struct {
	Ctx context.Context // In case you need a context.
}

type UserGetAllResult struct {
	Users []*entity.User `json:"users"`
	Error errs.Errs      `json:"error"`
}

type NewRoleV1Param struct {
	RoleRepo repository.Role
	Config   *Config
}

type RoleGetByUserIDParam struct {
	Ctx    context.Context
	UserId int `json:"user_id"`
}

type RoleGetByUserIDResult struct {
	Role  *entity.Role `json:"role"`
	Error errs.Errs    `json:"error"`
}
