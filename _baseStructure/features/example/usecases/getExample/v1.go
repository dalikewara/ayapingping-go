package getExample

import (
	"context"
	"github.com/dalikewara/ayapingping-go/v4/_baseStructure/domain"
	"github.com/dalikewara/ayapingping-go/v4/_baseStructure/features/example/utility"
)

type v1 struct {
	findExampleByID domain.FindExampleByIDRepository
}

func NewV1(findExampleByID domain.FindExampleByIDRepository) domain.GetExampleUseCase {
	return &v1{
		findExampleByID: findExampleByID,
	}
}

func (v *v1) ExecCtx(ctx context.Context, id uint64) (*domain.ExampleDTO1, error) {
	example, err := v.findExampleByID.ExecCtx(ctx, id)
	if err != nil {
		return nil, err
	}

	return utility.NewExampleDTO1(example), nil
}
