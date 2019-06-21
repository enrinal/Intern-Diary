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
	listcart, err := c.cart.FindCustCart(custid)
	if err != nil {
		return nil, err
	}
	return listcart, nil
}
