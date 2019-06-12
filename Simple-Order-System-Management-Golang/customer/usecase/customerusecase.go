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
	listcustomer, _ := c.customers.Fetch()
	return listcustomer
}

func (c *CustomerUsecase) GetCustomerById(ID int64) *models.Customer {
	customer, _ := c.customers.GetById(ID)
	return customer
}
