package usecase

import (
	"context"
	"time"

	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/customer"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

type CustomerUsecase struct {
	customers      customer.Repository
	contextTimeout time.Duration
}

func NewCustomerUsecase(customers customer.Repository, timeout time.Duration) customer.Usecase {
	return &CustomerUsecase{
		customers:      customers,
		contextTimeout: timeout,
	}
}

func (cust *CustomerUsecase) GetAllCustomer(c context.Context, num int64) ([]*models.Customer, error) {
	if num == 0 {
		num = 10
	}
	ctx, cancel := context.WithTimeout(c, cust.contextTimeout)
	defer cancel()

	listcustomer, err := cust.customers.Fetch(ctx, num)
	if err != nil {
		return nil, err
	}
	return listcustomer, nil
}

func (cust *CustomerUsecase) GetCustomerByID(c context.Context, ID int64) (*models.Customer, error) {
	ctx, cancel := context.WithTimeout(c, cust.contextTimeout)
	defer cancel()
	customer, err := cust.customers.GetCustomerById(ctx, ID)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
