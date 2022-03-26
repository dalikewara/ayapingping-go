package user_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dalikewara/ayapingping-go/src/domains/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNewMySQLRepository tests NewMySQLRepository function.
func TestNewMySQLRepository(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.Nil(t, err)
	repo := user.NewMySQLRepository(user.NewMySQLRepositoryParam{
		Db: db,
	})
	assert.Implements(t, (*user.RepositoryInterface)(nil), repo)
}

// TestMySQLRepository_FindAll tests MySQLRepository.FindAll method
// and all possible scenarios.
func TestMySQLRepository_FindAll(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.Nil(t, err)
	repo := user.NewMySQLRepository(user.NewMySQLRepositoryParam{
		Db: db,
	})
	ctx := context.Background()
	users := repo.FindAll(user.RepositoryFindAllParam{
		Ctx: ctx,
	})
	assert.Nil(t, users.Error)
	assert.NotNil(t, users.Users)
	userId := 0
	for _, v := range *users.Users {
		userId = v.Id
	}
	assert.Equal(t, 1, userId)
}

// TestMySQLRepository_FindByUsername tests MySQLRepository.FindByUsername method
// and all possible scenarios.
func TestMySQLRepository_FindByUsername(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.Nil(t, err)
	repo := user.NewMySQLRepository(user.NewMySQLRepositoryParam{
		Db: db,
	})
	ctx := context.Background()
	userRow := repo.FindByUsername(user.RepositoryFindByUsernameParam{
		Username: "dalikewara",
		Ctx:      ctx,
	})
	assert.Nil(t, userRow.Error)
	assert.NotNil(t, userRow.User)
	assert.Equal(t, 1, userRow.User.Id)
}

// TestMySQLRepository_Insert tests MySQLRepository.Insert method
// and all possible scenarios.
func TestMySQLRepository_Insert(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.Nil(t, err)
	repo := user.NewMySQLRepository(user.NewMySQLRepositoryParam{
		Db: db,
	})
	ctx := context.Background()
	reply := repo.Insert(user.RepositoryInsertParam{
		Username: "dalikewara",
		Ctx:      ctx,
	})
	assert.Nil(t, reply.Error)
}
