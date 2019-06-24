package usecase

import (
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/cart"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

type CartUsecase struct {
	cart cart.Repository
}

func NewCartUsecase(cart cart.Repository) CartUsecase {
	return CartUsecase{
		cart: cart,
	}
}

func (c *CartUsecase) GetCustCart(custid int64) ([]*models.Cart, error) {
	listcart, err := c.cart.FindByCustomerId(custid)
	if err != nil {
		return nil, err
	}
	return listcart, nil
}
func (c *CartUsecase) AddItem(item models.Item, qty int64, idcart int64) error {
	err := c.cart.Add(item.Name, qty, idcart)
	if err != nil {
		return err
	}
	return nil
}

func (c *CartUsecase) RemoveItem(item models.Item, qty int64, idcart int64) error {
	err := c.cart.Remove(item.Name, qty, idcart)
	if err != nil {
		return err
	}
	return nil
}
