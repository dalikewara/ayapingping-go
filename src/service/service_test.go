package service_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dalikewara/ayapingping-go/src/repository"
	"github.com/dalikewara/ayapingping-go/src/service"
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
		systemUserRole := "admin"
		systemUserIds := []int{1, 2, 3}
		config := &service.Config{
			SystemUserRole: systemUserRole,
			SystemUserIds:  systemUserIds,
		}
		svc := service.New(service.NewParam{
			Repo:   repo,
			Config: config,
		})
		assert.ObjectsAreEqual(service.Service{}, svc)
		assert.Implements(t, (*service.User)(nil), svc.User)
		assert.Implements(t, (*service.Role)(nil), svc.Role)
	})
}
