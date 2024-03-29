// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import models "gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// ChangeOrderDelivered provides a mock function with given fields: c, ID
func (_m *Usecase) ChangeOrderDelivered(c context.Context, ID int64) error {
	ret := _m.Called(c, ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(c, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChangeOrderProcess provides a mock function with given fields: c, ID
func (_m *Usecase) ChangeOrderProcess(c context.Context, ID int64) error {
	ret := _m.Called(c, ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(c, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChangeOrderSend provides a mock function with given fields: c, ID
func (_m *Usecase) ChangeOrderSend(c context.Context, ID int64) error {
	ret := _m.Called(c, ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(c, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckLimitOrder provides a mock function with given fields: c, ID
func (_m *Usecase) CheckLimitOrder(c context.Context, ID int64) (bool, error) {
	ret := _m.Called(c, ID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, int64) bool); ok {
		r0 = rf(c, ID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(c, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountOrderCust provides a mock function with given fields: c, ID
func (_m *Usecase) CountOrderCust(c context.Context, ID int64) (int, error) {
	ret := _m.Called(c, ID)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, int64) int); ok {
		r0 = rf(c, ID)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(c, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllOrder provides a mock function with given fields: c, num
func (_m *Usecase) GetAllOrder(c context.Context, num int64) ([]*models.Order, error) {
	ret := _m.Called(c, num)

	var r0 []*models.Order
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*models.Order); ok {
		r0 = rf(c, num)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(c, num)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllOrderById provides a mock function with given fields: c, ID
func (_m *Usecase) GetAllOrderById(c context.Context, ID int64) ([]*models.Order, error) {
	ret := _m.Called(c, ID)

	var r0 []*models.Order
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*models.Order); ok {
		r0 = rf(c, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(c, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderById provides a mock function with given fields: c, ID
func (_m *Usecase) GetOrderById(c context.Context, ID int64) (*models.Order, error) {
	ret := _m.Called(c, ID)

	var r0 *models.Order
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.Order); ok {
		r0 = rf(c, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(c, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
