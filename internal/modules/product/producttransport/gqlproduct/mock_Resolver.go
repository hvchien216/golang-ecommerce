// Code generated by mockery v2.20.0. DO NOT EDIT.

package gqlproduct

import (
	context "context"

	productmodel "github.com/hvchien216/golang-ecommerce/internal/modules/product/productmodel"
	mock "github.com/stretchr/testify/mock"
)

// MockResolver is an autogenerated mock type for the Resolver type
type MockResolver struct {
	mock.Mock
}

// CreateProduct provides a mock function with given fields: _a0, _a1
func (_m *MockResolver) CreateProduct(_a0 context.Context, _a1 productmodel.NewProductInput) (*productmodel.Product, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *productmodel.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, productmodel.NewProductInput) (*productmodel.Product, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, productmodel.NewProductInput) *productmodel.Product); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*productmodel.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, productmodel.NewProductInput) error); ok {
		r1 = rf(_a0, _a1)
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
