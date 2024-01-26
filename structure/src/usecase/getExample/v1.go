package getExample

import (
	"context"
	"errors"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/entity/example"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/repository/findExampleByID"
)

type v1 struct {
	findExampleByID findExampleByID.Contract
}

// ExecCtx executes the use case main logic with context
func (v *v1) ExecCtx(ctx context.Context, id uint64) (*example.DTO1, error) {
	exampleModel, err := v.findExampleByID.ExecCtx(ctx, id)
	if err != nil {
		return nil, err
	}
	if exampleModel == nil {
		return nil, errors.New("data not found")
	}

	return example.NewDTO1(exampleModel), nil
}

// NewV1 creates new v1 use case
func NewV1(findExampleByID findExampleByID.Contract) Contract {
	return &v1{
		findExampleByID: findExampleByID,
	}
}
