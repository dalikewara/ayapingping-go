package repository_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dalikewara/ayapingping-go/v2/src/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNewRoleMySQL tests NewRoleMySQL function.
func TestNewRoleMySQL(t *testing.T) {
	t.Run("Test OK", func(t *testing.T) {
		db, _, err := sqlmock.New()
		assert.Nil(t, err)
		repo := repository.NewRoleMySQL(repository.NewRoleMySQLParam{
			DB: db,
		})
		assert.Implements(t, (*repository.Role)(nil), repo)
	})
}

// TestRoleMySQL_FindByUserID tests roleMySQL.FindByUserID method
// and all possible scenarios.
func TestRoleMySQL_FindByUserID(t *testing.T) {
	db, _, _ := sqlmock.New()
	repo := repository.NewRoleMySQL(repository.NewRoleMySQLParam{
		DB: db,
	})
	ctx := context.Background()

	t.Run("ERR no role data found", func(t *testing.T) {
		expectedError := repository.ErrorRoleNoDataFound
		role := repo.FindByUserID(repository.RoleFindByUserIDParam{
			Ctx:    ctx,
			UserId: 2,
		})
		assert.Nil(t, role.Role)
		assert.NotNil(t, role.Error)
		assert.Equal(t, expectedError, role.Error)
	})

	t.Run("OK", func(t *testing.T) {
		role := repo.FindByUserID(repository.RoleFindByUserIDParam{
			Ctx:    ctx,
			UserId: 1,
		})
		assert.Nil(t, role.Error)
		assert.NotNil(t, role.Role)
		assert.Equal(t, 1, role.Role.UserId)
		assert.Equal(t, "admin", role.Role.Role)
	})
}
