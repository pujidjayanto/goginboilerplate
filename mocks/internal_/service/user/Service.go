// Code generated by mockery v2.50.4. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/pujidjayanto/goginboilerplate/internal/dto"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

type Service_Expecter struct {
	mock *mock.Mock
}

func (_m *Service) EXPECT() *Service_Expecter {
	return &Service_Expecter{mock: &_m.Mock}
}

// Login provides a mock function with given fields: _a0, _a1
func (_m *Service) Login(_a0 context.Context, _a1 dto.LoginRequest) (*dto.LoginResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 *dto.LoginResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.LoginRequest) (*dto.LoginResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.LoginRequest) *dto.LoginResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.LoginResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.LoginRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type Service_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 dto.LoginRequest
func (_e *Service_Expecter) Login(_a0 interface{}, _a1 interface{}) *Service_Login_Call {
	return &Service_Login_Call{Call: _e.mock.On("Login", _a0, _a1)}
}

func (_c *Service_Login_Call) Run(run func(_a0 context.Context, _a1 dto.LoginRequest)) *Service_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dto.LoginRequest))
	})
	return _c
}

func (_c *Service_Login_Call) Return(_a0 *dto.LoginResponse, _a1 error) *Service_Login_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_Login_Call) RunAndReturn(run func(context.Context, dto.LoginRequest) (*dto.LoginResponse, error)) *Service_Login_Call {
	_c.Call.Return(run)
	return _c
}

// Register provides a mock function with given fields: _a0, _a1
func (_m *Service) Register(_a0 context.Context, _a1 dto.RegisterRequest) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.RegisterRequest) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type Service_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 dto.RegisterRequest
func (_e *Service_Expecter) Register(_a0 interface{}, _a1 interface{}) *Service_Register_Call {
	return &Service_Register_Call{Call: _e.mock.On("Register", _a0, _a1)}
}

func (_c *Service_Register_Call) Run(run func(_a0 context.Context, _a1 dto.RegisterRequest)) *Service_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dto.RegisterRequest))
	})
	return _c
}

func (_c *Service_Register_Call) Return(_a0 error) *Service_Register_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_Register_Call) RunAndReturn(run func(context.Context, dto.RegisterRequest) error) *Service_Register_Call {
	_c.Call.Return(run)
	return _c
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
