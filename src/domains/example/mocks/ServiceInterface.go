// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	example "github.com/dalikewara/ayapingping-go/src/domains/example"
	mock "github.com/stretchr/testify/mock"
)

// ServiceInterface is an autogenerated mock type for the ServiceInterface type
type ServiceInterface struct {
	mock.Mock
}

// Get provides a mock function with given fields: param
func (_m *ServiceInterface) Get(param example.ServiceGetParam) example.ServiceGetResult {
	ret := _m.Called(param)

	var r0 example.ServiceGetResult
	if rf, ok := ret.Get(0).(func(example.ServiceGetParam) example.ServiceGetResult); ok {
		r0 = rf(param)
	} else {
		r0 = ret.Get(0).(example.ServiceGetResult)
	}

	return r0
}

// UpdateName provides a mock function with given fields: param
func (_m *ServiceInterface) UpdateName(param example.ServiceUpdateNameParam) example.ServiceUpdateNameResult {
	ret := _m.Called(param)

	var r0 example.ServiceUpdateNameResult
	if rf, ok := ret.Get(0).(func(example.ServiceUpdateNameParam) example.ServiceUpdateNameResult); ok {
		r0 = rf(param)
	} else {
		r0 = ret.Get(0).(example.ServiceUpdateNameResult)
	}

	return r0
}
