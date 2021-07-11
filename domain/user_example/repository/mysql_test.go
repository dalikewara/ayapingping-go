package repository_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dalikewara/ayapingping-go/domain/user_example"
	"github.com/dalikewara/ayapingping-go/domain/user_example/repository"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestNewMySQL tests `repository.NewMySQL` method.
func TestNewMySQL(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.Nil(t, err)
	repo := repository.NewMySQL(db)
	assert.Implements(t, (*user_example.Repository)(nil), repo)
}

// TestMySQL_GetByUsernamePassword tests method `GetByUsernamePassword` from `repository.MySQL`.
func TestMySQL_GetByUsernamePassword(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.Nil(t, err)
	repo := repository.NewMySQL(db)
	request := user_example.RepositoryGetByUsernamePasswordRequest{
		Username: "guest",
		Password: "guest",
	}
	response := repo.GetByUsernamePassword(request)
	assert.Nil(t, response.Error)
	assert.Equal(t, "guest", response.User.Username)
	assert.Equal(t, "guest",response.User.Password)
}

// TestMySQL_GetByUsernamePasswordContext tests method `GetByUsernamePasswordContext` from `repository.MySQL`.
func TestMySQL_GetByUsernamePasswordContext(t *testing.T) {
	ctx := context.Background()
	ctxWithTimeout, cancelFunc := context.WithTimeout(ctx, 1*time.Second)
	defer cancelFunc()
	db, _, err := sqlmock.New()
	assert.Nil(t, err)
	repo := repository.NewMySQL(db)
	request := user_example.RepositoryGetByUsernamePasswordContextRequest{
		Ctx: ctxWithTimeout,
		Username: "guest",
		Password: "guest",
	}
	response := repo.GetByUsernamePasswordContext(request)
	assert.Nil(t, response.User)
	assert.Error(t, response.Error)
}
