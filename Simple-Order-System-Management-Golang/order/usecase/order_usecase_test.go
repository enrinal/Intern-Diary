package usecase

import (
	"testing"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/order/mocks"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

func TestGetAllOrder (t *testing.T){
	t.Run("success", func(t *testing.T) {
		Ordermocks := &mocks.Usecase{}
		Ordermocks.On("GetAllOrder").Return([]*models.Order{
		&models.Order{ID : 1, Item:"Mobil"},
		&models.Order{ID : 1, Item:"Mobil"},
		&models.Order{ID : 1, Item:"Mobil"},
		&models.Order{ID : 1, Item:"Mobil"},
		&models.Order{ID : 1, Item:"Mobil"},
		&models.Order{ID : 1, Item:"Mobil"},
		&models.Order{ID : 1, Item:"Mobil"},
		&models.Order{ID : 1, Item:"Mobil"},
		&models.Order{ID : 1, Item:"Mobil"},
		&models.Order{ID : 1, Item:"Mobil"},
		}, nil)
		Order := NewOrderUsecase(Ordermocks)
		c := Order.GetAllOrder()
		if len(c) != 10{		
			t.Errorf("Error harus 10, yang di dapat %d",len(c))
		}
	})
}