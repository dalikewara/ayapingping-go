package repository

import (
	"github.com/dalikewara/ayapingping-go/v3/structure/src/adapter"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/config"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/repository/findUserDetailByID"
)

type Repository struct {
	FindUserDetailByID findUserDetailByID.Contract
}

// InitRepository initializes repositories
func InitRepository(cfg *config.Config, adapterClient *adapter.Adapter) (*Repository, error) {
	return &Repository{
		FindUserDetailByID: findUserDetailByID.NewMySQLDB(adapterClient.MySQLDB),
	}, nil
}
