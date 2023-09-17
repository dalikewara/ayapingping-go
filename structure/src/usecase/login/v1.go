package login

import (
	"context"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/entity/example"
)

type v1 struct{}

func (v *v1) ExecuteCtx(ctx context.Context, id string) (*example.BaseModel, error) {
	panic("implement me")
}

// New

func NewV1() Contract {
	return &v1{}
}
