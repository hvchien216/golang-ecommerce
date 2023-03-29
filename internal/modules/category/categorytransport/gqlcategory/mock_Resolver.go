// Code generated by mockery v2.20.0. DO NOT EDIT.

package gqlcategory

import (
	common "github.com/hvchien216/golang-ecommerce/common"
	categorymodel "github.com/hvchien216/golang-ecommerce/internal/modules/category/categorymodel"

	context "context"

	mock "github.com/stretchr/testify/mock"

	orm "github.com/hvchien216/golang-ecommerce/internal/orm"
)

// MockResolver is an autogenerated mock type for the Resolver type
type MockResolver struct {
	mock.Mock
}

// Categories provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockResolver) Categories(_a0 context.Context, _a1 *categorymodel.Filter, _a2 *common.Paging) (*categorymodel.CategoryResult, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *categorymodel.CategoryResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *categorymodel.Filter, *common.Paging) (*categorymodel.CategoryResult, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *categorymodel.Filter, *common.Paging) *categorymodel.CategoryResult); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*categorymodel.CategoryResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *categorymodel.Filter, *common.Paging) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Category provides a mock function with given fields: ctx, id
func (_m *MockResolver) Category(ctx context.Context, id int) (*orm.Category, error) {
	ret := _m.Called(ctx, id)

	var r0 *orm.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*orm.Category, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *orm.Category); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*orm.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Children provides a mock function with given fields: ctx, obj
func (_m *MockResolver) Children(ctx context.Context, obj *orm.Category) ([]*orm.Category, error) {
	ret := _m.Called(ctx, obj)

	var r0 []*orm.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *orm.Category) ([]*orm.Category, error)); ok {
		return rf(ctx, obj)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *orm.Category) []*orm.Category); ok {
		r0 = rf(ctx, obj)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*orm.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *orm.Category) error); ok {
		r1 = rf(ctx, obj)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCategory provides a mock function with given fields: _a0, _a1
func (_m *MockResolver) CreateCategory(_a0 context.Context, _a1 categorymodel.NewCategoryInput) (*orm.Category, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *orm.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, categorymodel.NewCategoryInput) (*orm.Category, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, categorymodel.NewCategoryInput) *orm.Category); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*orm.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, categorymodel.NewCategoryInput) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Parent provides a mock function with given fields: ctx, obj
func (_m *MockResolver) Parent(ctx context.Context, obj *orm.Category) (*orm.Category, error) {
	ret := _m.Called(ctx, obj)

	var r0 *orm.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *orm.Category) (*orm.Category, error)); ok {
		return rf(ctx, obj)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *orm.Category) *orm.Category); ok {
		r0 = rf(ctx, obj)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*orm.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *orm.Category) error); ok {
		r1 = rf(ctx, obj)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCategory provides a mock function with given fields: ctx, id, name, description, parentID
func (_m *MockResolver) UpdateCategory(ctx context.Context, id int, name string, description *string, parentID *int) (*orm.Category, error) {
	ret := _m.Called(ctx, id, name, description, parentID)

	var r0 *orm.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, *string, *int) (*orm.Category, error)); ok {
		return rf(ctx, id, name, description, parentID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, string, *string, *int) *orm.Category); ok {
		r0 = rf(ctx, id, name, description, parentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*orm.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, string, *string, *int) error); ok {
		r1 = rf(ctx, id, name, description, parentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockResolver interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockResolver creates a new instance of MockResolver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockResolver(t mockConstructorTestingTNewMockResolver) *MockResolver {
	mock := &MockResolver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
