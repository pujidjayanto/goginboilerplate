// Code generated by mockery v2.50.4. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

type Controller_Expecter struct {
	mock *mock.Mock
}

func (_m *Controller) EXPECT() *Controller_Expecter {
	return &Controller_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0
func (_m *Controller) Create(_a0 *gin.Context) {
	_m.Called(_a0)
}

// Controller_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type Controller_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 *gin.Context
func (_e *Controller_Expecter) Create(_a0 interface{}) *Controller_Create_Call {
	return &Controller_Create_Call{Call: _e.mock.On("Create", _a0)}
}

func (_c *Controller_Create_Call) Run(run func(_a0 *gin.Context)) *Controller_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *Controller_Create_Call) Return() *Controller_Create_Call {
	_c.Call.Return()
	return _c
}

func (_c *Controller_Create_Call) RunAndReturn(run func(*gin.Context)) *Controller_Create_Call {
	_c.Run(run)
	return _c
}

// NewController creates a new instance of Controller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewController(t interface {
	mock.TestingT
	Cleanup(func())
}) *Controller {
	mock := &Controller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
