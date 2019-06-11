package usecase

import (
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/customer"
)

const (
	RegularBuyer = iota + 1
	SubcriptionBuyer
)

type CustomerUsecase struct{
	customers customer.Repository
}

func NewCustomerUsecase(customers customer.Repository) CustomerUsecase {
	return CustomerUsecase{
		customers : customers,
	}
}
func (c *CustomerUsecase)GetAllCustomer() []*models.Customer{
	listcustomer,_ := c.customers.Fetch()
	return listcustomer
}

func (c *CustomerUsecase)GetCustomerById(ID int64) *models.Customer{
	customer,_ := c.customers.GetById(ID)
	return customer
}

func (c *CustomerUsecase)IsCustomerSubricption(ID int64)bool{
	customer,_ := c.customers.GetById(ID)
	if customer.Status == SubcriptionBuyer {
		return true
	}
	return false
}