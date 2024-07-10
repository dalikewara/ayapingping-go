package domain

import (
	"context"
	"github.com/dalikewara/ayapingping-go/v4/_base_structure/common"
	"time"
)

type ExampleRepository interface {
	FindByIDCtx(ctx context.Context, id uint64) (*Example, error)
}

type ExampleUseCase interface {
	GetDetailCtx(ctx context.Context, id uint64) (*ExampleDTO1, error)
}

type ExampleHttpService interface {
	ExampleDetail(method string, endpoint string)
}

type Example struct {
	ID        uint64    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (e *Example) SetCreatedAtNow() {
	e.CreatedAt = common.TimeNowUTC()
}

func (e *Example) ValidateUsername() error {
	return common.ValidateUsername(e.Username)
}

type ExampleDTO1 struct {
	ID        uint64    `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

func NewExampleDTO1(example *Example) *ExampleDTO1 {
	return &ExampleDTO1{
		ID:        example.ID,
		Username:  example.Username,
		CreatedAt: example.CreatedAt,
	}
}
