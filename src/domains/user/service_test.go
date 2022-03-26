package user_test

import (
	"context"
	"github.com/dalikewara/ayapingping-go/src/domains/user"
	"github.com/dalikewara/ayapingping-go/src/domains/user/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestNewService tests NewService function.
func TestNewService(t *testing.T) {
	repo := &mocks.RepositoryInterface{}
	service := user.NewService(user.NewServiceParam{
		Repository: repo,
	})
	assert.Implements(t, (*user.ServiceInterface)(nil), service)
}

// TestService_GetAll tests Service.GetAll method and all possible scenarios.
func TestService_GetAll(t *testing.T) {
	ctx := context.Background()
	param := user.ServiceGetAllParam{
		Ctx: ctx,
	}
	repoParam := user.RepositoryFindAllParam{
		Ctx: ctx,
	}
	repoResult := user.RepositoryFindAllResult{
		Users: &[]user.User{
			{
				Id:        1,
				Username:  "dalikewara",
				CreatedAt: time.Now(),
			},
		},
		Error: nil,
	}
	repo := &mocks.RepositoryInterface{}
	service := user.NewService(user.NewServiceParam{
		Repository: repo,
	})
	repo.On("FindAll", repoParam).Return(repoResult).Once()
	users := service.GetAll(param)
	repo.AssertCalled(t, "FindAll", repoParam)
	repo.AssertExpectations(t)
	assert.Nil(t, users.Error)
	userId := 0
	for _, v := range *users.Users {
		userId = v.Id
	}
	assert.Equal(t, 1, userId)
}

// TestService_GetByUsername tests Service.GetByUsername method and all possible scenarios.
func TestService_GetByUsername(t *testing.T) {
	ctx := context.Background()
	param := user.ServiceGetByUsernameParam{
		Username: "dalikewara",
		Ctx:      ctx,
	}
	repoParam := user.RepositoryFindByUsernameParam{
		Username: "dalikewara",
		Ctx:      ctx,
	}
	repoResult := user.RepositoryFindByUsernameResult{
		User: &user.User{
			Id:        1,
			Username:  "dalikewara",
			CreatedAt: time.Now(),
		},
		Error: nil,
	}
	repo := &mocks.RepositoryInterface{}
	service := user.NewService(user.NewServiceParam{
		Repository: repo,
	})
	repo.On("FindByUsername", repoParam).Return(repoResult).Once()
	userRow := service.GetByUsername(param)
	repo.AssertCalled(t, "FindByUsername", repoParam)
	repo.AssertExpectations(t)
	assert.Nil(t, userRow.Error)
	assert.Equal(t, 1, userRow.User.Id)
}

// TestService_Create tests Service.Create method and all possible scenarios.
func TestService_Create(t *testing.T) {
	ctx := context.Background()
	param := user.ServiceCreateParam{
		Username: "dalikewara",
		Ctx:      ctx,
	}
	repoParam := user.RepositoryInsertParam{
		Username: "dalikewara",
		Ctx:      ctx,
	}
	repoResult := user.RepositoryInsertResult{
		Error: nil,
	}
	repo := &mocks.RepositoryInterface{}
	service := user.NewService(user.NewServiceParam{
		Repository: repo,
	})
	repo.On("Insert", repoParam).Return(repoResult).Once()
	reply := service.Create(param)
	repo.AssertCalled(t, "Insert", repoParam)
	repo.AssertExpectations(t)
	assert.Nil(t, reply.Error)
}
