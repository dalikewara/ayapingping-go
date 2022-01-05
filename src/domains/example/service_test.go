package example_test

import (
	"github.com/dalikewara/ayapingping-go/src/domains/example"
	"github.com/dalikewara/ayapingping-go/src/domains/example/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNewService tests NewService function.
func TestNewService(t *testing.T) {
	repo := &mocks.RepositoryInterface{}
	srv := example.NewService(example.NewServiceParam{
		Repository: repo,
	})
	assert.Implements(t, (*example.ServiceInterface)(nil), srv)
}

// TestService_Get tests Service.Get method.
func TestService_Get(t *testing.T) {
	param := example.ServiceGetParam{
		Id: int64(1),
	}
	repoParam := example.RepositoryFindByIdParam{
		Id: int64(1),
	}
	repoResult := example.RepositoryFindByIdResult{
		Example: &example.Example{
			Id: int64(1),
			Name: "John Doe",
		},
		Error: nil,
	}
	repo := &mocks.RepositoryInterface{}
	srv := example.NewService(example.NewServiceParam{
		Repository: repo,
	})
	repo.On("FindById", repoParam).Return(repoResult).Once()
	res := srv.Get(param)
	repo.AssertCalled(t, "FindById", repoParam)
	repo.AssertExpectations(t)
	assert.Nil(t, res.Error)
	assert.Equal(t, param.Id, res.Example.Id)
	assert.Equal(t, "John Doe", res.Example.Name)
}

// TestService_UpdateName tests Service.UpdateName method.
func TestService_UpdateName(t *testing.T) {
	param := example.ServiceUpdateNameParam{
		Id: int64(1),
		Name: "Smith",
	}
	repoParam := example.RepositoryUpdateNameByIdParam{
		Id: int64(1),
		Name: "Smith",
	}
	repoResult := example.RepositoryUpdateNameByIdResult{
		Error: nil,
	}
	repo := &mocks.RepositoryInterface{}
	srv := example.NewService(example.NewServiceParam{
		Repository: repo,
	})
	repo.On("UpdateNameById", repoParam).Return(repoResult).Once()
	res := srv.UpdateName(param)
	repo.AssertCalled(t, "UpdateNameById", repoParam)
	repo.AssertExpectations(t)
	assert.Nil(t, res.Error)
}
