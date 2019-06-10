// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Fetch provides a mock function with given fields:
func (_m *Repository) Fetch() ([]*models.Customer, error) {
	ret := _m.Called()

	var r0 []*models.Customer
	if rf, ok := ret.Get(0).(func() []*models.Customer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Customer)
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

// GetById provides a mock function with given fields: id
func (_m *Repository) GetById(id int64) (*models.Customer, error) {
	ret := _m.Called(id)

	var r0 *models.Customer
	if rf, ok := ret.Get(0).(func(int64) *models.Customer); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Customer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
