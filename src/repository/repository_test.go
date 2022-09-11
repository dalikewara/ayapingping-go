package repository_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dalikewara/ayapingping-go/src/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNew tests New function.
func TestNew(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		mysqlDB, _, err := sqlmock.New()
		assert.Nil(t, err)
		repo := repository.New(repository.NewParam{
			MySQLDB: mysqlDB,
		})
		assert.ObjectsAreEqual(repository.Repository{}, repo)
		assert.Implements(t, (*repository.User)(nil), repo.User)
		assert.Implements(t, (*repository.Role)(nil), repo.Role)
	})
}
