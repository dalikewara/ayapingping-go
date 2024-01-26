package findExampleByID

import (
	"context"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/entity/example"
)

type Contract interface {
	ExecCtx(ctx context.Context, id uint64) (*example.Model, error)
}
