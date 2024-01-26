package repository

import (
	"github.com/dalikewara/ayapingping-go/v3/structure/src/adapter"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/config"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/repository/findExampleByID"
)

type Repository struct {
	FindExampleByID findExampleByID.Contract
}

// InitRepository initializes repositories
func InitRepository(cfg *config.Config, adapterClient *adapter.Adapter) (*Repository, error) {
	return &Repository{
		FindExampleByID: findExampleByID.NewMySQLDB(adapterClient.MySQLDB),
	}, nil
}
