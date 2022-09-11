package service

import (
	"github.com/dalikewara/ayapingping-go/src/repository"
)

type roleV1 struct {
	roleRepo repository.Role
	config   *Config
}

// NewRoleV1 generates new roleV1 that implements Role.
func NewRoleV1(param NewRoleV1Param) Role {
	return &roleV1{
		roleRepo: param.RoleRepo,
		config:   param.Config,
	}
}

// GetByUserID gets role data by user id.
func (s *roleV1) GetByUserID(param RoleGetByUserIDParam) RoleGetByUserIDResult {
	var result RoleGetByUserIDResult

	role := s.roleRepo.FindByUserID(repository.RoleFindByUserIDParam{
		Ctx:    param.Ctx,
		UserId: param.UserId,
	})
	if role.Error != nil {
		result.Error = role.Error
		return result
	}

	if !s.validateSystemUserRole(role.Role.UserId, role.Role.Role) {
		result.Error = ErrorRoleSystemUserNotAdmin
		return result
	}

	result.Role = role.Role

	return result
}

// validateSystemUserRole validates role if the user id is a system user.
func (s *roleV1) validateSystemUserRole(userId int, role string) bool {
	var isSystemUser bool

	for _, v := range s.config.SystemUserIds {
		if v == userId {
			isSystemUser = true
			break
		}
	}

	if isSystemUser && role != s.config.SystemUserRole {
		return false
	}

	return true
}
