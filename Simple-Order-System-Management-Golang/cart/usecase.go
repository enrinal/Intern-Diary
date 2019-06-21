package cart

import (
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

type Usecase interface {
	GetCustCart(custid int64) ([]*models.Cart, error)
	AddItem(item models.Item, qty int64) error
	// RemoveItem(item string) error
}
