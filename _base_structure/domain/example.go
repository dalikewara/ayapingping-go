package domain

import (
	"context"
	"github.com/dalikewara/ayapingping-go/v4/_base_structure/common"
	"time"
)

type FindExampleByIDRepository interface {
	ExecCtx(ctx context.Context, id uint64) (*Example, error)
}

type GetExampleUseCase interface {
	ExecCtx(ctx context.Context, id uint64) (*ExampleDTO1, error)
}

type ExampleDelivery interface {
	RegisterHandler(method string, endpoint string)
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

func (e *Example) ToDTO1() *ExampleDTO1 {
	return &ExampleDTO1{
		ID:        e.ID,
		Username:  e.Username,
		CreatedAt: e.CreatedAt,
	}
}

type ExampleDTO1 struct {
	ID        uint64    `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

type ExampleJSONPresenter struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
}
