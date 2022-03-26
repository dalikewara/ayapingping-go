// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	user "github.com/dalikewara/ayapingping-go/src/domains/user"
	mock "github.com/stretchr/testify/mock"
)

// RepositoryInterface is an autogenerated mock type for the RepositoryInterface type
type RepositoryInterface struct {
	mock.Mock
}

// FindAll provides a mock function with given fields: param
func (_m *RepositoryInterface) FindAll(param user.RepositoryFindAllParam) user.RepositoryFindAllResult {
	ret := _m.Called(param)

	var r0 user.RepositoryFindAllResult
	if rf, ok := ret.Get(0).(func(user.RepositoryFindAllParam) user.RepositoryFindAllResult); ok {
		r0 = rf(param)
	} else {
		r0 = ret.Get(0).(user.RepositoryFindAllResult)
	}

	return r0
}

// FindByUsername provides a mock function with given fields: param
func (_m *RepositoryInterface) FindByUsername(param user.RepositoryFindByUsernameParam) user.RepositoryFindByUsernameResult {
	ret := _m.Called(param)

	var r0 user.RepositoryFindByUsernameResult
	if rf, ok := ret.Get(0).(func(user.RepositoryFindByUsernameParam) user.RepositoryFindByUsernameResult); ok {
		r0 = rf(param)
	} else {
		r0 = ret.Get(0).(user.RepositoryFindByUsernameResult)
	}

	return r0
}

// Insert provides a mock function with given fields: param
func (_m *RepositoryInterface) Insert(param user.RepositoryInsertParam) user.RepositoryInsertResult {
	ret := _m.Called(param)

	var r0 user.RepositoryInsertResult
	if rf, ok := ret.Get(0).(func(user.RepositoryInsertParam) user.RepositoryInsertResult); ok {
		r0 = rf(param)
	} else {
		r0 = ret.Get(0).(user.RepositoryInsertResult)
	}

	return r0
}
