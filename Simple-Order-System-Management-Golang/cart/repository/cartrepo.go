package repository

import (
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

func FindCustCart(idcust int64) ([]*models.Cart, error) {
	listcart := make([]*models.Cart, 0)
	return listcart, nil
}

func Add(item string, qty int64) error {
	return nil
}

func Remove(item string) error {
	return nil
}
