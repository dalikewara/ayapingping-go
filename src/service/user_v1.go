package service

import (
	"github.com/dalikewara/ayapingping-go/src/repository"
)

type userV1 struct {
	userRepo repository.User
}

// NewUserV1 generates new userV1 that implements User.
func NewUserV1(param NewUserV1Param) User {
	return &userV1{
		userRepo: param.UserRepo,
	}
}

// GetAll gets all user data.
func (s *userV1) GetAll(param UserGetAllParam) UserGetAllResult {
	var result UserGetAllResult

	users := s.userRepo.FindAll(repository.UserFindAllParam{
		Ctx: param.Ctx,
	})
	if users.Error != nil {
		result.Error = users.Error
		return result
	}

	result.Users = users.Users

	return result
}
