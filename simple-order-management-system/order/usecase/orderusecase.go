package usecase

import (
	"context"
	"time"

	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/customer"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/order"
)

const (
	RegularBuyer = iota + 1
	SubcriptionBuyer
)

const (
	Pending = iota + 1
	Process
	Send
	Delivered
)

type OrderUsecase struct {
	orders         order.Repository
	customer       customer.Repository
	contextTimeout time.Duration
}

func NewOrderUsecase(orders order.Repository, customer customer.Repository, timeout time.Duration) order.Usecase {
	return &OrderUsecase{
		orders:         orders,
		customer:       customer,
		contextTimeout: timeout,
	}
}

func (o *OrderUsecase) GetAllOrder(c context.Context, num int64) ([]*models.Order, error) {
	ctx, cancel := context.WithTimeout(c, o.contextTimeout)
	defer cancel()
	listorder, err := o.orders.GetAllOrder(ctx, num)
	if err != nil {
		return nil, err
	}
	return listorder, nil
}

func (o *OrderUsecase) GetOrderById(c context.Context, ID int64) (*models.Order, error) {
	ctx, cancel := context.WithTimeout(c, o.contextTimeout)
	defer cancel()
	order, err := o.orders.GetOrderById(ctx, ID)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *OrderUsecase) ChangeOrderProcess(c context.Context, ID int64) error {
	ctx, cancel := context.WithTimeout(c, o.contextTimeout)
	defer cancel()
	err := o.orders.ChangeOrderProcess(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderUsecase) ChangeOrderSend(c context.Context, ID int64) error {
	ctx, cancel := context.WithTimeout(c, o.contextTimeout)
	defer cancel()
	err := o.orders.ChangeOrderSend(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderUsecase) ChangeOrderDelivered(c context.Context, ID int64) error {
	ctx, cancel := context.WithTimeout(c, o.contextTimeout)
	defer cancel()
	err := o.orders.ChangeOrderDelivered(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderUsecase) GetAllOrderById(c context.Context, ID int64) ([]*models.Order, error) {
	ctx, cancel := context.WithTimeout(c, o.contextTimeout)
	defer cancel()
	listorder, err := o.orders.GetAllOrderById(ctx, ID)
	if err != nil {
		return nil, err
	}
	return listorder, nil
}

func (o *OrderUsecase) CountOrderCust(c context.Context, ID int64) (int, error) {
	ctx, cancel := context.WithTimeout(c, o.contextTimeout)
	defer cancel()
	listorder, err := o.orders.GetAllOrderById(ctx, ID)
	if err != nil {
		return 0, err
	}
	return len(listorder), nil
}

func (o *OrderUsecase) CheckLimitOrder(c context.Context, ID int64) (bool, error) {
	ctx, cancel := context.WithTimeout(c, o.contextTimeout)
	defer cancel()
	customer, err := o.customer.GetCustomerById(ctx, ID)
	if err != nil {
		return false, err
	}
	order, err := o.orders.GetAllOrderById(ctx, ID)
	if err != nil {
		return false, err
	}
	if (customer.Status == RegularBuyer && len(order) <= 5) || (customer.Status == SubcriptionBuyer) {
		return true, nil
	}
	return false, nil
}
