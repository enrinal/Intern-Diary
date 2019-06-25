package cart

import (
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

type Repository interface {
	FindByCustomerId(custid int64) ([]*models.Cart, error)
	Add(cart models.Cart) error
	Remove(item string, qty int64, idcart int64) error
}
