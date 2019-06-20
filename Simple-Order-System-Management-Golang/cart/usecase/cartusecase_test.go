package usecase

import (
	"testing"

	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/cart/mocks"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

func TestGetCustCart(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		Cartmocks := &mocks.Usecase{}
		Cartmocks.On("GetCustCart").Return([]*models.Cart{
			&models.Cart{IDCart: 1, IDCust: 2, Items: "mobil"},
			&models.Cart{IDCart: 2, IDCust: 2, Items: "mobil"},
			&models.Cart{IDCart: 3, IDCust: 2, Items: "mobil"},
		}, nil)
		Cart := NewCartUsecase(Cartmocks)
		custid := 2
		c, _ := Cart.GetCustCart(int64(custid))
		if len(c) != 3 {
			t.Errorf("Harus mendapatakan 3, yang didapat %d", len(c))
		}
	})
}
