package service_test

import (
	"context"
	"fmt"
	"github.com/dalikewara/ayapingping-go/src/entity"
	"github.com/dalikewara/ayapingping-go/src/library/errs"
	"github.com/dalikewara/ayapingping-go/src/repository"
	"github.com/dalikewara/ayapingping-go/src/repository/mocks"
	"github.com/dalikewara/ayapingping-go/src/service"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestNewUserV1 tests NewUserV1 function.
func TestNewUserV1(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		userRepo := mocks.NewUser(t)
		svc := service.NewUserV1(service.NewUserV1Param{
			UserRepo: userRepo,
		})
		assert.Implements(t, (*service.User)(nil), svc)
	})
}

// TestUserV1_GetAll tests userV1.GetAll method
// and all possible scenarios.
func TestUserV1_GetAll(t *testing.T) {
	userRepo := mocks.NewUser(t)
	svc := service.NewUserV1(service.NewUserV1Param{
		UserRepo: userRepo,
	})

	t.Run("ERR userRepoFindAll", func(t *testing.T) {
		ctx := context.Background()
		expectedError := errs.New("TEST", "err")
		actualError := errs.New("TEST", "err")
		userRepoFindAllParam := repository.UserFindAllParam{
			Ctx: ctx,
		}
		userRepoFindAllResult := repository.UserFindAllResult{
			Error: actualError,
		}
		userRepo.On("FindAll", userRepoFindAllParam).Return(userRepoFindAllResult).Once()
		param := service.UserGetAllParam{
			Ctx: ctx,
		}
		users := svc.GetAll(param)
		userRepo.AssertCalled(t, "FindAll", userRepoFindAllParam)
		userRepo.AssertExpectations(t)
		assert.NotNil(t, users.Error)
		assert.Equal(t, expectedError, users.Error)
	})

	t.Run("OK", func(t *testing.T) {
		ctx := context.Background()
		var expectedUserRepoUsers, actualUserRepoUsers []*entity.User
		for i := 1; i <= 5; i++ {
			timeNow := time.Now()
			expectedUserRepoUsers = append(expectedUserRepoUsers, &entity.User{
				Id:        i,
				Username:  fmt.Sprintf("johndoe%v", i),
				CreatedAt: timeNow,
			})
			actualUserRepoUsers = append(actualUserRepoUsers, &entity.User{
				Id:        i,
				Username:  fmt.Sprintf("johndoe%v", i),
				CreatedAt: timeNow,
			})
		}
		userRepoFindAllParam := repository.UserFindAllParam{
			Ctx: ctx,
		}
		userRepoFindAllResult := repository.UserFindAllResult{
			Users: actualUserRepoUsers,
		}
		userRepo.On("FindAll", userRepoFindAllParam).Return(userRepoFindAllResult).Once()
		param := service.UserGetAllParam{
			Ctx: ctx,
		}
		users := svc.GetAll(param)
		userRepo.AssertCalled(t, "FindAll", userRepoFindAllParam)
		userRepo.AssertExpectations(t)
		assert.Nil(t, users.Error)
		for i := 0; i < len(users.Users); i++ {
			assert.Equal(t, expectedUserRepoUsers[i], users.Users[i])
		}
	})
}
