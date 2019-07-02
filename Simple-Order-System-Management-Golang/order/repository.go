package order

import (
	"context"

	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

type Repository interface {
	GetAllOrder(ctx context.Context, num int64) ([]*models.Order, error)
	GetAllOrderById(ctx context.Context, ID int64) ([]*models.Order, error)
	GetOrderById(ctx context.Context, ID int64) (*models.Order, error)
	ChangeOrderSend(ctx context.Context, ID int64) error
	ChangeOrderProcess(ID int64) error
	ChangeOrderDelivered(ID int64) error
}
