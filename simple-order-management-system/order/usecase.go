package order

import (
	"context"

	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

type Usecase interface {
	GetAllOrder(c context.Context, num int64) ([]*models.Order, error)
	GetAllOrderById(c context.Context, ID int64) ([]*models.Order, error)
	GetOrderById(c context.Context, ID int64) (*models.Order, error)
	ChangeOrderSend(c context.Context, ID int64) error
	ChangeOrderProcess(c context.Context, ID int64) error
	ChangeOrderDelivered(c context.Context, ID int64) error
	CheckLimitOrder(c context.Context, ID int64) (bool, error)
	CountOrderCust(c context.Context, ID int64) (int, error)
}
