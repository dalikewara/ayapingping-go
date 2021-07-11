package usecase

import (
	"github.com/dalikewara/ayapingping-go/domain/user_example"
	"github.com/dalikewara/ayapingping-go/domain/user_example/helpers/print_me"
)

type UseCase struct {
	repo user_example.Repository
}

// NewUseCase generates new `UseCase` use case that implements `user_example.UseCase`.
func NewUseCase(repo user_example.Repository) user_example.UseCase {
	return &UseCase{
		repo: repo,
	}
}

// Login logins user_example.
func (uc *UseCase) Login(request user_example.UseCaseLoginRequest) user_example.UseCaseLoginResponse {
	u := user_example.Entity{
		Username: request.Username,
		Password: request.Password,
	}
	repoResponse := uc.repo.GetByUsernamePassword(user_example.RepositoryGetByUsernamePasswordRequest{
		Username: u.Username,
		Password: u.Password,
	})
	print_me.PrintMe("Logged in: " + repoResponse.User.Username)
	return user_example.UseCaseLoginResponse{
		User: repoResponse.User,
		Error: nil,
	}
}

// LoginContext logins user_example with context.
func (uc *UseCase) LoginContext(request user_example.UseCaseLoginContextRequest) user_example.UseCaseLoginContextResponse {
	u := user_example.Entity{
		Username: request.Username,
		Password: request.Password,
	}
	repoResponse := uc.repo.GetByUsernamePasswordContext(user_example.RepositoryGetByUsernamePasswordContextRequest{
		Ctx: request.Ctx,
		Username: u.Username,
		Password: u.Password,
	})
	if repoResponse.Error != nil {
		return user_example.UseCaseLoginContextResponse{
			User: nil,
			Error: repoResponse.Error,
		}
	}
	print_me.PrintMe("Logged in: " + repoResponse.User.Username + ", but you should not see this")
	return user_example.UseCaseLoginContextResponse{
		User: repoResponse.User,
		Error: nil,
	}
}
