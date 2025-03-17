// Code generated by mockery v2.50.4. DO NOT EDIT.

package mocks

import (
	context "context"

	pagination "github.com/pujidjayanto/goginboilerplate/pkg/pagination"
	mock "github.com/stretchr/testify/mock"

	product "github.com/pujidjayanto/goginboilerplate/internal/repository/product"
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

// FindAll provides a mock function with given fields: _a0
func (_m *Repository) FindAll(_a0 context.Context) ([]*product.Product, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []*product.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*product.Product, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*product.Product); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*product.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type Repository_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *Repository_Expecter) FindAll(_a0 interface{}) *Repository_FindAll_Call {
	return &Repository_FindAll_Call{Call: _e.mock.On("FindAll", _a0)}
}

func (_c *Repository_FindAll_Call) Run(run func(_a0 context.Context)) *Repository_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Repository_FindAll_Call) Return(_a0 []*product.Product, _a1 error) *Repository_FindAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_FindAll_Call) RunAndReturn(run func(context.Context) ([]*product.Product, error)) *Repository_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// FindAllPaginated provides a mock function with given fields: _a0, _a1, _a2
func (_m *Repository) FindAllPaginated(_a0 context.Context, _a1 product.ProductFilter, _a2 pagination.PaginationRequest) ([]*product.Product, int64, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for FindAllPaginated")
	}

	var r0 []*product.Product
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, product.ProductFilter, pagination.PaginationRequest) ([]*product.Product, int64, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, product.ProductFilter, pagination.PaginationRequest) []*product.Product); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*product.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, product.ProductFilter, pagination.PaginationRequest) int64); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, product.ProductFilter, pagination.PaginationRequest) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Repository_FindAllPaginated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAllPaginated'
type Repository_FindAllPaginated_Call struct {
	*mock.Call
}

// FindAllPaginated is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 product.ProductFilter
//   - _a2 pagination.PaginationRequest
func (_e *Repository_Expecter) FindAllPaginated(_a0 interface{}, _a1 interface{}, _a2 interface{}) *Repository_FindAllPaginated_Call {
	return &Repository_FindAllPaginated_Call{Call: _e.mock.On("FindAllPaginated", _a0, _a1, _a2)}
}

func (_c *Repository_FindAllPaginated_Call) Run(run func(_a0 context.Context, _a1 product.ProductFilter, _a2 pagination.PaginationRequest)) *Repository_FindAllPaginated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(product.ProductFilter), args[2].(pagination.PaginationRequest))
	})
	return _c
}

func (_c *Repository_FindAllPaginated_Call) Return(_a0 []*product.Product, _a1 int64, _a2 error) *Repository_FindAllPaginated_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *Repository_FindAllPaginated_Call) RunAndReturn(run func(context.Context, product.ProductFilter, pagination.PaginationRequest) ([]*product.Product, int64, error)) *Repository_FindAllPaginated_Call {
	_c.Call.Return(run)
	return _c
}

// FindById provides a mock function with given fields: _a0, _a1
func (_m *Repository) FindById(_a0 context.Context, _a1 uint) (*product.Product, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *product.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) (*product.Product, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint) *product.Product); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*product.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_FindById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindById'
type Repository_FindById_Call struct {
	*mock.Call
}

// FindById is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uint
func (_e *Repository_Expecter) FindById(_a0 interface{}, _a1 interface{}) *Repository_FindById_Call {
	return &Repository_FindById_Call{Call: _e.mock.On("FindById", _a0, _a1)}
}

func (_c *Repository_FindById_Call) Run(run func(_a0 context.Context, _a1 uint)) *Repository_FindById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint))
	})
	return _c
}

func (_c *Repository_FindById_Call) Return(_a0 *product.Product, _a1 error) *Repository_FindById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_FindById_Call) RunAndReturn(run func(context.Context, uint) (*product.Product, error)) *Repository_FindById_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *Repository) Update(_a0 context.Context, _a1 *product.Product) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *product.Product) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type Repository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *product.Product
func (_e *Repository_Expecter) Update(_a0 interface{}, _a1 interface{}) *Repository_Update_Call {
	return &Repository_Update_Call{Call: _e.mock.On("Update", _a0, _a1)}
}

func (_c *Repository_Update_Call) Run(run func(_a0 context.Context, _a1 *product.Product)) *Repository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*product.Product))
	})
	return _c
}

func (_c *Repository_Update_Call) Return(_a0 error) *Repository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_Update_Call) RunAndReturn(run func(context.Context, *product.Product) error) *Repository_Update_Call {
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
