package example_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dalikewara/ayapingping-go/src/domains/example"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNewMySQLRepository tests `example.NewMySQLRepository` method.
func TestNewMySQLRepository(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.Nil(t, err)
	repo := example.NewMySQLRepository(example.NewMySQLRepositoryParam{
		Db: db,
	})
	assert.Implements(t, (*example.RepositoryInterface)(nil), repo)
}

// TestMySQLRepository_FindById tests `MySQLRepository.FindById` method.
func TestMySQLRepository_FindById(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.Nil(t, err)
	repo := example.NewMySQLRepository(example.NewMySQLRepositoryParam{
		Db: db,
	})
	res := repo.FindById(example.RepositoryFindByIdParam{
		Id: int64(1),
	})
	assert.Nil(t, res.Error)
	assert.Equal(t, int64(1), res.Example.Id)
}

// TestMySQLRepository_UpdateNameById tests `MySQLRepository.UpdateNameById` method.
func TestMySQLRepository_UpdateNameById(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.Nil(t, err)
	repo := example.NewMySQLRepository(example.NewMySQLRepositoryParam{
		Db: db,
	})
	res := repo.UpdateNameById(example.RepositoryUpdateNameByIdParam{
		Id: 1,
		Name: "Smith",
	})
	assert.Nil(t, res.Error)
}
