// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import models "gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
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
