package usecase

import (
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/customer"
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
