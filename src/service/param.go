package service

import (
	"context"
	"github.com/dalikewara/ayapingping-go/v2/src/entity"
	"github.com/dalikewara/ayapingping-go/v2/src/repository"
)

type Config struct {
	Example string `json:"example"`
}

type NewParam struct {
	Repo   *repository.Repository
	Config *Config
}

type NewExampleParam struct {
	ExampleRepo repository.Example
}

type ExampleGetAllParam struct {
	Ctx context.Context // In case you need a context.
}

type ExampleGetAllResult struct {
	Examples []*entity.Example `json:"examples"`
	Error    error             `json:"error"`
}
