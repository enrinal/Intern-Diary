package cart

import (
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

type Repository interface {
	FindByCustomerId(idcust int64) ([]models.Cart, error)
	Add(cart models.Cart) error
	Remove(cart models.Cart) error
}
