package customer

import (
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

// Repository represent the customer repository contract
type Repository interface {
	Fetch() ([]*models.Customer, error)
	GetById(id int64) (*models.Customer, error)
}
