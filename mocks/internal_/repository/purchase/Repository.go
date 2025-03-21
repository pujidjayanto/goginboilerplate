// Code generated by mockery v2.50.4. DO NOT EDIT.

package mocks

import (
	context "context"

	purchase "github.com/pujidjayanto/goginboilerplate/internal/repository/purchase"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

type Repository_Expecter struct {
	mock *mock.Mock
}

func (_m *Repository) EXPECT() *Repository_Expecter {
	return &Repository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *Repository) Create(_a0 context.Context, _a1 purchase.Purchase) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, purchase.Purchase) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type Repository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 purchase.Purchase
func (_e *Repository_Expecter) Create(_a0 interface{}, _a1 interface{}) *Repository_Create_Call {
	return &Repository_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *Repository_Create_Call) Run(run func(_a0 context.Context, _a1 purchase.Purchase)) *Repository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(purchase.Purchase))
	})
	return _c
}

func (_c *Repository_Create_Call) Return(_a0 error) *Repository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_Create_Call) RunAndReturn(run func(context.Context, purchase.Purchase) error) *Repository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
