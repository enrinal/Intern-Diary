package customer

import (
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

// Usecase represent the customer Usecase contract
type Usecase interface {
	GetAllCustomer() []*models.Customer
	GetCustomerById(ID int64) (*models.Customer, error)
}
