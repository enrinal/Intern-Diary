package usecase

import (
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/customer"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

type CustomerUsecase struct {
	customers customer.Repository
}

func NewCustomerUsecase(customers customer.Repository) CustomerUsecase {
	return CustomerUsecase{
		customers: customers,
	}
}
func (c *CustomerUsecase) GetAllCustomer() []*models.Customer {
	listcustomer, err := c.customers.Fetch()
	if err != nil {
		return nil
	}
	return listcustomer
}

func (c *CustomerUsecase) GetCustomerById(ID int64) (*models.Customer, error) {
	customer, err := c.customers.GetCustomerById(ID)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
