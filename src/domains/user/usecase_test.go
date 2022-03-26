package user_test

import (
	"context"
	"github.com/dalikewara/ayapingping-go/src/domains/user"
	"github.com/dalikewara/ayapingping-go/src/domains/user/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestNewUseCase tests NewUseCase function.
func TestNewUseCase(t *testing.T) {
	service := &mocks.ServiceInterface{}
	useCase := user.NewUseCase(user.NewUseCaseParam{
		Service: service,
	})
	assert.Implements(t, (*user.UseCaseInterface)(nil), useCase)
}

// TestUseCase_GetAll tests UseCase.GetAll method and all possible scenarios.
func TestUseCase_GetAll(t *testing.T) {
	ctx := context.Background()
	param := user.UseCaseGetAllParam{
		Ctx: ctx,
	}
	serviceParam := user.ServiceGetAllParam{
		Ctx: ctx,
	}
	serviceResult := user.ServiceGetAllResult{
		Users: &[]user.User{
			{
				Id:        1,
				Username:  "dalikewara",
				CreatedAt: time.Now(),
			},
		},
		Error: nil,
	}
	service := &mocks.ServiceInterface{}
	useCase := user.NewUseCase(user.NewUseCaseParam{
		Service: service,
	})
	service.On("GetAll", serviceParam).Return(serviceResult).Once()
	users := useCase.GetAll(param)
	service.AssertCalled(t, "GetAll", serviceParam)
	service.AssertExpectations(t)
	assert.Nil(t, users.Error)
	userId := 0
	for _, v := range *users.Users {
		userId = v.Id
	}
	assert.Equal(t, 1, userId)
}

// TestUseCase_Create tests UseCase.Create method and all possible scenarios.
func TestUseCase_Create(t *testing.T) {
	ctx := context.Background()
	param := user.UseCaseCreateParam{
		Username: "dalikewara2",
		Ctx:      ctx,
	}
	serviceGetByUsernameParam := user.ServiceGetByUsernameParam{
		Username: "dalikewara2",
		Ctx:      ctx,
	}
	serviceGetByUsernameResult := user.ServiceGetByUsernameResult{
		User:  nil,
		Error: nil,
	}
	serviceCreateParam := user.ServiceCreateParam{
		Username: "dalikewara2",
		Ctx:      ctx,
	}
	serviceCreateResult := user.ServiceCreateResult{
		Error: nil,
	}
	service := &mocks.ServiceInterface{}
	useCase := user.NewUseCase(user.NewUseCaseParam{
		Service: service,
	})
	service.On("GetByUsername", serviceGetByUsernameParam).Return(serviceGetByUsernameResult).Once()
	service.On("Create", serviceCreateParam).Return(serviceCreateResult).Once()
	reply := useCase.Create(param)
	service.AssertCalled(t, "Create", serviceCreateParam)
	service.AssertCalled(t, "GetByUsername", serviceGetByUsernameParam)
	service.AssertExpectations(t)
	assert.Nil(t, reply.Error)
}
