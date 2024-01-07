package getUserDetail

import (
	"context"
	"errors"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/entity/user"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/entity/userProfile"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/repository/findUserDetailByID"
)

type v1 struct {
	findUserDetailByID findUserDetailByID.Contract
}

// ExecCtx executes the use case main logic with context
func (v *v1) ExecCtx(ctx context.Context, userID uint64) (*user.DTO1, error) {
	userModel, err := v.findUserDetailByID.ExecCtx(ctx, userID)
	if err != nil {
		return nil, err
	}
	if userModel == nil {
		return nil, errors.New("user not found")
	}

	return user.NewDTO1(&userModel.Model, userProfile.NewDTO1(userModel.Profile)), nil
}

// NewV1 creates new v1 use case
func NewV1(findUserDetailByID findUserDetailByID.Contract) Contract {
	return &v1{
		findUserDetailByID: findUserDetailByID,
	}
}
