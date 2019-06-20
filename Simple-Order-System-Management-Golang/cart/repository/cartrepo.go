package repository

import (
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

func GetCustCart() ([]*models.Cart, error) {
	listcart := make([]*models.Cart, 0)
	return listcart, nil
}

func AddItem(item string, qty int64) error {
	return nil
}

func RemoveItem(item string) error {
	return nil
}
