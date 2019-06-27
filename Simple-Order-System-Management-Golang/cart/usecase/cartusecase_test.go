package usecase

import (
	"testing"

	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/cart/mocks"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

func TestGetCustCart(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		Cartmocks := &mocks.Repository{}
		Cartmocks.On("FindByCustomerId", int64(2)).Return([]models.Cart{
			models.Cart{IDCart: 1, IDCust: 2, Items: []models.Item{models.Item{Id: int64(1), Name: "mobil"},
				models.Item{Id: int64(1), Name: "mobil"}}},
			models.Cart{IDCart: 1, IDCust: 2, Items: []models.Item{models.Item{Id: int64(1), Name: "mobil"},
				models.Item{Id: int64(1), Name: "mobil"}}},
			models.Cart{IDCart: 1, IDCust: 2, Items: []models.Item{models.Item{Id: int64(1), Name: "mobil"},
				models.Item{Id: int64(1), Name: "mobil"}}},
		}, nil)
		cart := NewCartUsecase(Cartmocks)
		c, _ := cart.GetCustCart(int64(2))
		if len(c) != 3 {
			t.Errorf("Harus mendapatakan 3, yang didapat %d", len(c))
		}
	})
}

func TestAddItem(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		cartmodel := models.Cart{IDCart: 1000, IDCust: 2, Items: []models.Item{models.Item{Id: int64(1), Name: "mobil"},
			models.Item{Id: int64(1), Name: "mobil"}}}
		Cartmocks := &mocks.Repository{}
		Cartmocks.On("Add", cartmodel).Return(nil)
		cart := NewCartUsecase(Cartmocks)
		err := cart.AddItem(cartmodel)
		if err != nil {
			print(err)
		}
	})
}

func TestRemoveItem(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		cartmodel := models.Cart{IDCart: 1000, IDCust: 2, Items: []models.Item{models.Item{Id: int64(1), Name: "mobil"},
			models.Item{Id: int64(1), Name: "mobil"}}}
		Cartmocks := &mocks.Repository{}
		Cartmocks.On("Remove", cartmodel).Return(nil)
		cart := NewCartUsecase(Cartmocks)
		err := cart.RemoveItem(cartmodel)
		if err != nil {
			print(err)
		}
	})
}
