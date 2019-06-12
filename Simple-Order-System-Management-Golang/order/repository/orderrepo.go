package repository

import (
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)


func GetAllOrder() ([]*models.Order, error) {
	listorder := make([]*models.Order, 0)
	return listorder, nil
}

func GetOrderById(ID int64) (*models.Order, error) {
	order := &models.Order{}
	return order, nil
}

func GetAllOrderById(ID int64) ([]*models.Order, error) {
	listorder := make([]*models.Order, 0)
	return listorder, nil
}