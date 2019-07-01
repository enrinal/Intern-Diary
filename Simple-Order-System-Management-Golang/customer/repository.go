package customer

import (
	"context"

	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

// Repository represent the customer repository contract
type Repository interface {
	Fetch(ctx context.Context, num int64) (result []*models.Customer, err error)
	GetCustomerById(ctx context.Context, id int64) (*models.Customer, error)
}
