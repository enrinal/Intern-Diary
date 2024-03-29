// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"

import mock "github.com/stretchr/testify/mock"
import models "gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, cust
func (_m *Repository) Add(ctx context.Context, cust *models.Customer) error {
	ret := _m.Called(ctx, cust)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Customer) error); ok {
		r0 = rf(ctx, cust)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Repository) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx, num
func (_m *Repository) Fetch(ctx context.Context, num int64) ([]*models.Customer, error) {
	ret := _m.Called(ctx, num)

	var r0 []*models.Customer
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*models.Customer); ok {
		r0 = rf(ctx, num)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Customer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, num)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCustomerById provides a mock function with given fields: ctx, id
func (_m *Repository) GetCustomerById(ctx context.Context, id int64) (*models.Customer, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Customer
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.Customer); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Customer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, cust
func (_m *Repository) Update(ctx context.Context, cust *models.Customer) error {
	ret := _m.Called(ctx, cust)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Customer) error); ok {
		r0 = rf(ctx, cust)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
