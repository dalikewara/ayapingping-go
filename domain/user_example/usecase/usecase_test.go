package usecase_test

import (
	"context"
	"github.com/dalikewara/ayapingping-go/domain/user_example"
	"github.com/dalikewara/ayapingping-go/domain/user_example/mocks"
	"github.com/dalikewara/ayapingping-go/domain/user_example/usecase"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestNewUseCase tests `usecase.NewUseCase` method.
func TestNewUseCase(t *testing.T) {
	repo := &mocks.Repository{}
	uc := usecase.NewUseCase(repo)
	assert.Implements(t, (*user_example.UseCase)(nil), uc)
}

// TestUseCase_Login tests method `Login` from `usecase.UseCase`.
func TestUseCase_Login(t *testing.T) {
	request := user_example.UseCaseLoginRequest{
		Username: "guest",
		Password: "guest",
	}
	repoRequest := user_example.RepositoryGetByUsernamePasswordRequest{
		Username: request.Username,
		Password: request.Password,
	}
	repoResponse := user_example.RepositoryGetByUsernamePasswordResponse{
		User: &user_example.Entity{
			Username: repoRequest.Username,
			Password: repoRequest.Password,
		},
		Error: nil,
	}
	repo := &mocks.Repository{}
	uc := usecase.NewUseCase(repo)
	repo.On("GetByUsernamePassword", repoRequest).Return(repoResponse).Once()
	response := uc.Login(request)
	repo.AssertCalled(t, "GetByUsernamePassword", repoRequest)
	repo.AssertExpectations(t)
	assert.Nil(t, response.Error)
	assert.Equal(t, request.Username, response.User.Username)
	assert.Equal(t, request.Password, response.User.Password)
}

// TestUseCase_LoginContext tests method `LoginContext` from `usecase.UseCase`.
func TestUseCase_LoginContext(t *testing.T) {
	t.Run("context ok", func(t *testing.T) {
		ctx := context.Background()
		request := user_example.UseCaseLoginContextRequest{
			Ctx:      ctx,
			Username: "guest",
			Password: "guest",
		}
		repoRequest := user_example.RepositoryGetByUsernamePasswordContextRequest{
			Ctx:      request.Ctx,
			Username: request.Username,
			Password: request.Password,
		}
		repoResponse := user_example.RepositoryGetByUsernamePasswordContextResponse{
			User: &user_example.Entity{
				Username: repoRequest.Username,
				Password: repoRequest.Password,
			},
			Error: nil,
		}
		repo := &mocks.Repository{}
		uc := usecase.NewUseCase(repo)
		repo.On("GetByUsernamePasswordContext", repoRequest).Return(repoResponse).Once()
		response := uc.LoginContext(request)
		repo.AssertCalled(t, "GetByUsernamePasswordContext", repoRequest)
		repo.AssertExpectations(t)
		assert.Nil(t, response.Error)
		assert.Equal(t, request.Username, response.User.Username)
		assert.Equal(t, request.Password, response.User.Password)
	})
	t.Run("context error", func(t *testing.T) {
		ctx := context.Background()
		ctxWithTimeout, cancelFunc := context.WithTimeout(ctx, 1*time.Second)
		defer cancelFunc()
		time.Sleep(2 * time.Second)
		request := user_example.UseCaseLoginContextRequest{
			Ctx:      ctxWithTimeout,
			Username: "guest",
			Password: "guest",
		}
		repoRequest := user_example.RepositoryGetByUsernamePasswordContextRequest{
			Ctx:      request.Ctx,
			Username: request.Username,
			Password: request.Password,
		}
		repoResponse := user_example.RepositoryGetByUsernamePasswordContextResponse{
			User: nil,
			Error: nil,
		}
		select {
			case <-repoRequest.Ctx.Done():
				repoResponse.Error = repoRequest.Ctx.Err()
		}
		repo := &mocks.Repository{}
		uc := usecase.NewUseCase(repo)
		repo.On("GetByUsernamePasswordContext", repoRequest).Return(repoResponse).Once()
		response := uc.LoginContext(request)
		repo.AssertCalled(t, "GetByUsernamePasswordContext", repoRequest)
		repo.AssertExpectations(t)
		assert.Nil(t, response.User)
		assert.Error(t, response.Error)
	})
}
