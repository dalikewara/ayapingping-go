package example

import (
	"context"
	"github.com/dalikewara/ayapingping-go/v4/_base_structure/domain"
)

type repositoryMySQL struct {
	client interface{}
}

func NewRepositoryMySQL(client interface{}) domain.ExampleRepository {
	return &repositoryMySQL{
		client: client,
	}
}

func (r *repositoryMySQL) FindByIDCtx(ctx context.Context, id uint64) (*domain.Example, error) {
	example := &domain.Example{
		ID:       id,
		Username: "dalikewara",
		Password: "admin123",
	}

	example.SetCreatedAtNow()

	return example, nil
}
