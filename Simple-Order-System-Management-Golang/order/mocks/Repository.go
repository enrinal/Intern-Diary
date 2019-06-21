// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// ChangeOrderDelivered provides a mock function with given fields: ID
func (_m *Repository) ChangeOrderDelivered(ID int64) error {
	ret := _m.Called(ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChangeOrderProcess provides a mock function with given fields: ID
func (_m *Repository) ChangeOrderProcess(ID int64) error {
	ret := _m.Called(ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChangeOrderSend provides a mock function with given fields: ID
func (_m *Repository) ChangeOrderSend(ID int64) error {
	ret := _m.Called(ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllOrder provides a mock function with given fields:
func (_m *Repository) GetAllOrder() ([]*models.Order, error) {
	ret := _m.Called()

	var r0 []*models.Order
	if rf, ok := ret.Get(0).(func() []*models.Order); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllOrderById provides a mock function with given fields: ID
func (_m *Repository) GetAllOrderById(ID int64) ([]*models.Order, error) {
	ret := _m.Called(ID)

	var r0 []*models.Order
	if rf, ok := ret.Get(0).(func(int64) []*models.Order); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderById provides a mock function with given fields: ID
func (_m *Repository) GetOrderById(ID int64) (*models.Order, error) {
	ret := _m.Called(ID)

	var r0 *models.Order
	if rf, ok := ret.Get(0).(func(int64) *models.Order); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}