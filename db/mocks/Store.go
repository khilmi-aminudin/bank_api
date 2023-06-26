// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	repositories "github.com/khilmi-aminudin/bank_api/repositories"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// CreateCustomer provides a mock function with given fields: ctx, arg
func (_m *Store) CreateCustomer(ctx context.Context, arg repositories.CreateCustomerParams) (repositories.MCustomer, error) {
	ret := _m.Called(ctx, arg)

	var r0 repositories.MCustomer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, repositories.CreateCustomerParams) (repositories.MCustomer, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, repositories.CreateCustomerParams) repositories.MCustomer); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(repositories.MCustomer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, repositories.CreateCustomerParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllCustomers provides a mock function with given fields: ctx
func (_m *Store) GetAllCustomers(ctx context.Context) ([]repositories.MCustomer, error) {
	ret := _m.Called(ctx)

	var r0 []repositories.MCustomer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]repositories.MCustomer, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []repositories.MCustomer); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repositories.MCustomer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCustomerByEmail provides a mock function with given fields: ctx, email
func (_m *Store) GetCustomerByEmail(ctx context.Context, email string) (repositories.GetCustomerByEmailRow, error) {
	ret := _m.Called(ctx, email)

	var r0 repositories.GetCustomerByEmailRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (repositories.GetCustomerByEmailRow, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) repositories.GetCustomerByEmailRow); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(repositories.GetCustomerByEmailRow)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCustomerById provides a mock function with given fields: ctx, id
func (_m *Store) GetCustomerById(ctx context.Context, id uuid.UUID) (repositories.MCustomer, error) {
	ret := _m.Called(ctx, id)

	var r0 repositories.MCustomer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (repositories.MCustomer, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) repositories.MCustomer); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(repositories.MCustomer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCustomerByUsername provides a mock function with given fields: ctx, username
func (_m *Store) GetCustomerByUsername(ctx context.Context, username string) (repositories.GetCustomerByUsernameRow, error) {
	ret := _m.Called(ctx, username)

	var r0 repositories.GetCustomerByUsernameRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (repositories.GetCustomerByUsernameRow, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) repositories.GetCustomerByUsernameRow); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(repositories.GetCustomerByUsernameRow)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewStore creates a new instance of Store. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStore(t mockConstructorTestingTNewStore) *Store {
	mock := &Store{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
