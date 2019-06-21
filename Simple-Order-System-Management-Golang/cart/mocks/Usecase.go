// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// GetCustCart provides a mock function with given fields: custid
func (_m *Usecase) GetCustCart(custid int64) ([]*models.Cart, error) {
	ret := _m.Called(custid)

	var r0 []*models.Cart
	if rf, ok := ret.Get(0).(func(int64) []*models.Cart); ok {
		r0 = rf(custid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Cart)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(custid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
