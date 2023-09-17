package login

import (
	"context"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/entity/example"
)

type Contract interface {
	ExecuteCtx(ctx context.Context, id string) (*example.BaseModel, error)
}
