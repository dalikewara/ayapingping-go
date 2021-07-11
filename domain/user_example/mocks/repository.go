// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	user_example "github.com/dalikewara/ayapingping-go/domain/user_example"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetByUsernamePassword provides a mock function with given fields: request
func (_m *Repository) GetByUsernamePassword(request user_example.RepositoryGetByUsernamePasswordRequest) user_example.RepositoryGetByUsernamePasswordResponse {
	ret := _m.Called(request)

	var r0 user_example.RepositoryGetByUsernamePasswordResponse
	if rf, ok := ret.Get(0).(func(user_example.RepositoryGetByUsernamePasswordRequest) user_example.RepositoryGetByUsernamePasswordResponse); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(user_example.RepositoryGetByUsernamePasswordResponse)
	}

	return r0
}

// GetByUsernamePasswordContext provides a mock function with given fields: request
func (_m *Repository) GetByUsernamePasswordContext(request user_example.RepositoryGetByUsernamePasswordContextRequest) user_example.RepositoryGetByUsernamePasswordContextResponse {
	ret := _m.Called(request)

	var r0 user_example.RepositoryGetByUsernamePasswordContextResponse
	if rf, ok := ret.Get(0).(func(user_example.RepositoryGetByUsernamePasswordContextRequest) user_example.RepositoryGetByUsernamePasswordContextResponse); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(user_example.RepositoryGetByUsernamePasswordContextResponse)
	}

	return r0
}
