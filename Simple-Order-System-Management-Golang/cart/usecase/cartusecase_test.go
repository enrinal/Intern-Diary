package usecase

import (
	"testing"

	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/cart/mocks"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

func TestGetCustCart(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		Cartmocks := &mocks.Usecase{}
		Cartmocks.On("GetCustCart", int64(2)).Return([]*models.Cart{
			&models.Cart{IDCart: 1, IDCust: 2, Items: []string{"mobil", "motor"}},
			&models.Cart{IDCart: 2, IDCust: 2, Items: []string{"mobil", "motor"}},
			&models.Cart{IDCart: 3, IDCust: 2, Items: []string{"mobil", "motor"}},
		}, nil)
		cart := NewCartUsecase(Cartmocks)
		c, _ := cart.GetCustCart(int64(2))
		if len(c) != 3 {
			t.Errorf("Harus mendapatakan 3, yang didapat %d", len(c))
		}
	})
}
