package usecase

import (
	"github.com/dalikewara/ayapingping-go/v3/structure/src/config"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/repository"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/usecase/getExample"
)

type UseCase struct {
	GetExample getExample.Contract
}

// InitUseCase initializes use case
func InitUseCase(cfg *config.Config, repo *repository.Repository) (*UseCase, error) {
	return &UseCase{
		GetExample: getExample.NewV1(repo.FindExampleByID),
	}, nil
}
