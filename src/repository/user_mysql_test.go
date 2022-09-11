package repository_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dalikewara/ayapingping-go/v2/src/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNewUserMySQL tests NewUserMySQL function.
func TestNewUserMySQL(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		db, _, err := sqlmock.New()
		assert.Nil(t, err)
		repo := repository.NewUserMySQL(repository.NewUserMySQLParam{
			DB: db,
		})
		assert.Implements(t, (*repository.User)(nil), repo)
	})
}

// TestUserMySQL_FindAll tests userMySQL.FindAll method
// and all possible scenarios.
func TestUserMySQL_FindAll(t *testing.T) {
	db, _, _ := sqlmock.New()
	repo := repository.NewUserMySQL(repository.NewUserMySQLParam{
		DB: db,
	})
	ctx := context.Background()

	t.Run("OK", func(t *testing.T) {
		users := repo.FindAll(repository.UserFindAllParam{
			Ctx: ctx,
		})
		assert.Nil(t, users.Error)
		assert.NotNil(t, users.Users)
		assert.Equal(t, 1, users.Users[0].Id)
	})
}
