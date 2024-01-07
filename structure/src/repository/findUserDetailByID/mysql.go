package findUserDetailByID

import (
	"context"
	"database/sql"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/entity/user"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/entity/userProfile"
)

type mysqlDB struct {
	db *sql.DB
}

// ExecCtx executes the repository main logic with context
func (s *mysqlDB) ExecCtx(ctx context.Context, id uint64) (*user.ModelWithDetail, error) {
	// Communicates with the real MySQL connection here, for example: querying.
	userModel := user.Model{
		ID:       id,
		Username: "dalikewara",
	}

	userModel.SetCreatedAtNow()

	profileModel := userProfile.Model{
		ID:     1,
		UserID: userModel.ID,
		Name:   "Dali Kewara",
	}

	profileModel.SetCreatedAtNow()

	return &user.ModelWithDetail{
		Model:   userModel,
		Profile: &profileModel,
	}, nil
}

// NewMySQLDB creates new MySQL database repository
func NewMySQLDB(db *sql.DB) Contract {
	return &mysqlDB{
		db: db,
	}
}
