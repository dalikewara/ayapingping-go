package service

type User interface {
	// GetAll gets all user data.
	GetAll(param UserGetAllParam) UserGetAllResult
}

type Role interface {
	// GetByUserID gets role data by user id.
	GetByUserID(param RoleGetByUserIDParam) RoleGetByUserIDResult
}

type Service struct {
	User User
	Role Role
}

// New generates new service.
func New(param NewParam) *Service {
	userV1Service := NewUserV1(NewUserV1Param{
		UserRepo: param.Repo.User,
	})

	roleV1Service := NewRoleV1(NewRoleV1Param{
		RoleRepo: param.Repo.Role,
		Config:   param.Config,
	})

	return &Service{
		User: userV1Service,
		Role: roleV1Service,
	}
}
