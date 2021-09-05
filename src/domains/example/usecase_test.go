package example_test

import (
	"github.com/dalikewara/ayapingping-go/src/domains/example"
	"github.com/dalikewara/ayapingping-go/src/domains/example/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNewUseCase tests `example.NewUseCase` method.
func TestNewUseCase(t *testing.T) {
	srv := &mocks.ServiceInterface{}
	uc := example.NewUseCase(example.NewUseCaseParam{
		ExampleService: srv,
	})
	assert.Implements(t, (*example.UseCaseInterface)(nil), uc)
}

// TestUseCase_GetAndChangeName tests `UseCase.GetAndChangeName` method.
func TestUseCase_GetAndChangeName(t *testing.T) {
	param := example.UseCaseGetAndChangeNameParam{
		Id: int64(1),
		Name: "Smith",
	}
	srvGetParam := example.ServiceGetParam{
		Id: int64(1),
	}
	srvUpdateNameParam := example.ServiceUpdateNameParam{
		Id: int64(1),
		Name: "Smith",
	}
	srvGetResult := example.ServiceGetResult{
		Example: &example.Example{
			Id: int64(1),
			Name: "John Doe",
		},
		Error: nil,
	}
	srvUpdateNameResult := example.ServiceUpdateNameResult{
		Error: nil,
	}
	srv := &mocks.ServiceInterface{}
	uc := example.NewUseCase(example.NewUseCaseParam{
		ExampleService: srv,
	})
	srv.On("Get", srvGetParam).Return(srvGetResult).Once()
	srv.On("UpdateName", srvUpdateNameParam).Return(srvUpdateNameResult).Once()
	res := uc.GetAndChangeName(example.UseCaseGetAndChangeNameParam{
		Id: int64(1),
		Name: "Smith",
	})
	srv.AssertCalled(t, "Get", srvGetParam)
	srv.AssertCalled(t, "UpdateName", srvUpdateNameParam)
	srv.AssertExpectations(t)
	assert.Nil(t, res.Error)
	assert.Equal(t, param.Id, res.Example.Id)
	assert.Equal(t, "Smith", res.Example.Name)
}
