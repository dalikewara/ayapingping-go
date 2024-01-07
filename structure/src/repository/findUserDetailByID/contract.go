package findUserDetailByID

import (
	"context"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/entity/user"
)

type Contract interface {
	ExecCtx(ctx context.Context, id uint64) (*user.ModelWithDetail, error)
}
