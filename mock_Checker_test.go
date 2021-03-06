package main

import "github.com/stretchr/testify/mock"

import "github.com/gambol99/go-marathon"

type MockChecker struct {
	mock.Mock
}

// Name provides a mock function with given fields:
func (_m *MockChecker) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Check provides a mock function with given fields: _a0
func (_m *MockChecker) Check(_a0 marathon.Application) AppCheck {
	ret := _m.Called(_a0)

	var r0 AppCheck
	if rf, ok := ret.Get(0).(func(marathon.Application) AppCheck); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(AppCheck)
	}

	return r0
}
