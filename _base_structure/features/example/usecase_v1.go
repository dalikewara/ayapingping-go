package example

import (
	"context"
	"github.com/dalikewara/ayapingping-go/v4/_base_structure/domain"
)

type useCaseV1 struct {
	exampleRepository domain.ExampleRepository
}

func NewUseCaseV1(exampleRepository domain.ExampleRepository) domain.ExampleUseCase {
	return &useCaseV1{
		exampleRepository: exampleRepository,
	}
}

func (u *useCaseV1) GetDetailCtx(ctx context.Context, id uint64) (*domain.ExampleDTO1, error) {
	example, err := u.exampleRepository.FindByIDCtx(ctx, id)
	if err != nil {
		return nil, err
	}

	return domain.NewExampleDTO1(example), nil
}
