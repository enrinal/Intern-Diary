package usecase

import (
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/order"
)

const (
	Pending = iota + 1
	Process
	Send
	Delivered
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

func (o *OrderUsecase) ChangeOrderProcess(ID int64) string {
	order,_ := o.orders.GetOrderById(ID)
	if order.Status == Pending {
		order.Status = Process
		return "Order Process"
	}
	return "Error Change Status to Process"
}

func (o *OrderUsecase) ChangeOrderSend(ID int64) string {
	order,_ := o.orders.GetOrderById(ID)
	if order.Status == Process {
		order.Status = Send
		return "Order Send"
	}
	return "Error Change Status to Send"
}

func (o *OrderUsecase) ChangeOrderDelivered(ID int64) string {
	order,_ := o.orders.GetOrderById(ID)
	if order.Status == Send {
		order.Status = Delivered
		return "Order Delivered"
	}
	return "Error Change Status to Delivered"
}

func (o *OrderUsecase) GetAllOrderById(ID int64) []*models.Order{
	listorder,_ := o.orders.GetAllOrderById(ID)
	return listorder
}