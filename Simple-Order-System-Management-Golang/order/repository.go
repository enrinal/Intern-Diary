package order

import (
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

type Repository interface {
	GetAllOrder() ([]*models.Order, error)
	GetAllOrderById(ID int64) ([]*models.Order, error)
	GetOrderById(ID int64) (*models.Order, error)
	ChangeOrderSend(ID int64) (error)
	ChangeOrderProcess(ID int64) (error)
	ChangeOrderDelivered(ID int64) (error)
	GetOrderStatus(ID int64)(status string)
}