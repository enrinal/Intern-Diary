package cart

import (
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

type Usecase interface {
	GetCustCart(custid int64) ([]models.Cart, error)
	AddItem(cart models.Cart) error
	RemoveItem(cart models.Cart) error
}
