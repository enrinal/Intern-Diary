package customer

import (
	"context"

	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

// Usecase represent the customer Usecase contract
type Usecase interface {
	GetAllCustomer(ctx context.Context, num int64) ([]*models.Customer, error)
	GetCustomerByID(ctx context.Context, id int64) (*models.Customer, error)
	Update(ctx context.Context, cust *models.Customer) error
	Add(context.Context, *models.Customer) error
	Delete(ctx context.Context, id int64) error
}
