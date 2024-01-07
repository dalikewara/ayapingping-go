package getUserDetail

import (
	"context"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/entity/user"
)

type Contract interface {
	ExecCtx(ctx context.Context, userID uint64) (*user.DTO1, error)
}
