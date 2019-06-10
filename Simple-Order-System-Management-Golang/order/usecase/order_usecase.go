package usecase

import (
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/order"
)

type OrderUsecase struct{
	orders order.Repository
}

func NewOrderUsecase(orders order.Repository) OrderUsecase {
	return OrderUsecase{
		orders : orders,
	}
}

func (o *OrderUsecase)GetAllOrder() []*models.Order{
	listorder,_ := o.orders.GetAllOrder()
	return listorder
}

func (o *OrderUsecase)GetOrderById(ID int64) *models.Order{
	order,_ := o.orders.GetOrderById(ID)
	return order
}