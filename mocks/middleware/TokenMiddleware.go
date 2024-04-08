// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// TokenMiddleware is an autogenerated mock type for the TokenMiddleware type
type TokenMiddleware struct {
	mock.Mock
}

// Employee provides a mock function with given fields: next
func (_m *TokenMiddleware) Employee(next echo.HandlerFunc) echo.HandlerFunc {
	ret := _m.Called(next)

	if len(ret) == 0 {
		panic("no return value specified for Employee")
	}

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func(echo.HandlerFunc) echo.HandlerFunc); ok {
		r0 = rf(next)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// NewTokenMiddleware creates a new instance of TokenMiddleware. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTokenMiddleware(t interface {
	mock.TestingT
	Cleanup(func())
}) *TokenMiddleware {
	mock := &TokenMiddleware{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}