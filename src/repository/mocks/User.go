// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	repository "github.com/dalikewara/ayapingping-go/v2/src/repository"
	mock "github.com/stretchr/testify/mock"
)

// User is an autogenerated mock type for the User type
type User struct {
	mock.Mock
}

// FindAll provides a mock function with given fields: param
func (_m *User) FindAll(param repository.UserFindAllParam) repository.UserFindAllResult {
	ret := _m.Called(param)

	var r0 repository.UserFindAllResult
	if rf, ok := ret.Get(0).(func(repository.UserFindAllParam) repository.UserFindAllResult); ok {
		r0 = rf(param)
	} else {
		r0 = ret.Get(0).(repository.UserFindAllResult)
	}

	return r0
}

type mockConstructorTestingTNewUser interface {
	mock.TestingT
	Cleanup(func())
}

// NewUser creates a new instance of User. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUser(t mockConstructorTestingTNewUser) *User {
	mock := &User{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}