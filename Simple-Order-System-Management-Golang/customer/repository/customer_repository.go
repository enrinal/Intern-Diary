package repository

import (
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

func Fetch() ([]*models.Customer, error) {

	listcustomer := make([]*models.Customer, 0)
	return listcustomer, nil
}

func GetById(id int) (*models.Customer, error) {
	customer := &models.Customer{}
	return customer, nil
}
