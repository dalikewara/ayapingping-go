package usecase

import (
	"github.com/dalikewara/ayapingping-go/v3/structure/src/config"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/repository"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/usecase/getUserDetail"
)

type UseCase struct {
	GetUserDetail getUserDetail.Contract
}

// InitUseCase initializes use case
func InitUseCase(cfg *config.Config, repo *repository.Repository) (*UseCase, error) {
	return &UseCase{
		GetUserDetail: getUserDetail.NewV1(repo.FindUserDetailByID),
	}, nil
}
