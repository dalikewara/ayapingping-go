package service_test

import (
	"context"
	"github.com/dalikewara/ayapingping-go/v2/src/entity"
	"github.com/dalikewara/ayapingping-go/v2/src/library/errs"
	"github.com/dalikewara/ayapingping-go/v2/src/repository"
	"github.com/dalikewara/ayapingping-go/v2/src/repository/mocks"
	"github.com/dalikewara/ayapingping-go/v2/src/service"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestNewRoleV1 tests NewRoleV1 function.
func TestNewRoleV1(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		roleRepo := mocks.NewRole(t)
		svc := service.NewRoleV1(service.NewRoleV1Param{
			RoleRepo: roleRepo,
		})
		assert.Implements(t, (*service.Role)(nil), svc)
	})
}

// TestRoleV1_GetByUserID tests roleV1.GetByUserID method
// and all possible scenarios.
func TestRoleV1_GetByUserID(t *testing.T) {
	roleRepo := mocks.NewRole(t)
	systemUserRole := "admin"
	systemUserIds := []int{1, 2, 3}
	config := &service.Config{
		SystemUserRole: systemUserRole,
		SystemUserIds:  systemUserIds,
	}
	svc := service.NewRoleV1(service.NewRoleV1Param{
		RoleRepo: roleRepo,
		Config:   config,
	})

	t.Run("ERR roleRepoFindByUserID", func(t *testing.T) {
		ctx := context.Background()
		expectedError := errs.New("TEST", "err")
		actualError := errs.New("TEST", "err")
		roleRepoFindByUserIDParam := repository.RoleFindByUserIDParam{
			Ctx: ctx,
		}
		roleRepoFindByUserIDResult := repository.RoleFindByUserIDResult{
			Error: actualError,
		}
		roleRepo.On("FindByUserID", roleRepoFindByUserIDParam).Return(roleRepoFindByUserIDResult).Once()
		param := service.RoleGetByUserIDParam{
			Ctx: ctx,
		}
		roles := svc.GetByUserID(param)
		roleRepo.AssertCalled(t, "FindByUserID", roleRepoFindByUserIDParam)
		roleRepo.AssertExpectations(t)
		assert.NotNil(t, roles.Error)
		assert.Equal(t, expectedError, roles.Error)
	})

	t.Run("ERR the user is a system user but the role is not an admin", func(t *testing.T) {
		ctx := context.Background()
		expectedError := service.ErrorRoleSystemUserNotAdmin
		actualRoleRepoRole := &entity.Role{
			Id:        1,
			UserId:    1,
			Role:      "basic",
			CreatedAt: time.Now(),
		}
		roleRepoFindByUserIDParam := repository.RoleFindByUserIDParam{
			Ctx:    ctx,
			UserId: 1,
		}
		roleRepoFindByUserIDResult := repository.RoleFindByUserIDResult{
			Role: actualRoleRepoRole,
		}
		roleRepo.On("FindByUserID", roleRepoFindByUserIDParam).Return(roleRepoFindByUserIDResult).Once()
		param := service.RoleGetByUserIDParam{
			Ctx:    ctx,
			UserId: 1,
		}
		roles := svc.GetByUserID(param)
		roleRepo.AssertCalled(t, "FindByUserID", roleRepoFindByUserIDParam)
		roleRepo.AssertExpectations(t)
		assert.NotNil(t, roles.Error)
		assert.Equal(t, expectedError, roles.Error)
	})

	t.Run("OK", func(t *testing.T) {
		ctx := context.Background()
		timeNow := time.Now()
		expectedRoleRepoRole := &entity.Role{
			Id:        1,
			UserId:    1,
			Role:      "admin",
			CreatedAt: timeNow,
		}
		actualRoleRepoRole := &entity.Role{
			Id:        1,
			UserId:    1,
			Role:      "admin",
			CreatedAt: timeNow,
		}
		roleRepoFindByUserIDParam := repository.RoleFindByUserIDParam{
			Ctx:    ctx,
			UserId: 1,
		}
		roleRepoFindByUserIDResult := repository.RoleFindByUserIDResult{
			Role: actualRoleRepoRole,
		}
		roleRepo.On("FindByUserID", roleRepoFindByUserIDParam).Return(roleRepoFindByUserIDResult).Once()
		param := service.RoleGetByUserIDParam{
			Ctx:    ctx,
			UserId: 1,
		}
		role := svc.GetByUserID(param)
		roleRepo.AssertCalled(t, "FindByUserID", roleRepoFindByUserIDParam)
		roleRepo.AssertExpectations(t)
		assert.Nil(t, role.Error)
		assert.Equal(t, expectedRoleRepoRole, role.Role)
	})
}
