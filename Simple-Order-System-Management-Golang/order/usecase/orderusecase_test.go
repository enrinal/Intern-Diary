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

func TestGetOrderById (t *testing.T){
	t.Run("success", func(t *testing.T) {
		Ordermocks := &mocks.Usecase{}
		Ordermocks.On("GetOrderById", int64(2)).Return(&models.Order{ID : 2, Item:"Mobil", Status:1}, nil)
		Order := NewOrderUsecase(Ordermocks)
		c := Order.GetOrderById(int64(2))

		if c == nil {
			t.Errorf("Error not found id")
		}
	})
}

func TestOrderProcess(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		Ordermocks := &mocks.Usecase{}
		Ordermocks.On("GetOrderById", int64(2)).Return(&models.Order{ID : 2, Item:"Mobil", Status:1}, nil)
		Order := NewOrderUsecase(Ordermocks)
		c := Order.ChangeOrderProcess(int64(2))
		if c != "Order Process" {
			t.Errorf("Error not found id")
		}
	})
}

func TestOrderSend(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		Ordermocks := &mocks.Usecase{}
		Ordermocks.On("GetOrderById", int64(2)).Return(&models.Order{ID : 2, Item:"Mobil", Status:2}, nil)
		Order := NewOrderUsecase(Ordermocks)
		c := Order.ChangeOrderSend(int64(2))
		if c != "Order Send" {
			t.Errorf("Error not found id")
		}
	})
}

func TestOrderDelivered(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		Ordermocks := &mocks.Usecase{}
		Ordermocks.On("GetOrderById", int64(2)).Return(&models.Order{ID : 2, Item:"Mobil", Status:3}, nil)
		Order := NewOrderUsecase(Ordermocks)
		c := Order.ChangeOrderDelivered(int64(2))
		if c != "Order Delivered" {
			t.Errorf("Error not found id")
		}
	})
}

func TestAllOrderByCustomerId (t *testing.T){
	t.Run("success", func(t *testing.T) {
		Ordermocks := &mocks.Usecase{}
		Ordermocks.On("GetAllOrderById",int64(1)).Return([]*models.Order{
		&models.Order{IDCust : 1, Item:"Mobil"},
		&models.Order{IDCust : 1, Item:"Mobil"},
		&models.Order{IDCust : 1, Item:"Mobil"},
		&models.Order{IDCust : 1, Item:"Mobil"},
		&models.Order{IDCust : 1, Item:"Mobil"},
		&models.Order{IDCust : 1, Item:"Mobil"},
		&models.Order{IDCust : 1, Item:"Mobil"},
		&models.Order{IDCust : 1, Item:"Mobil"},
		&models.Order{IDCust : 1, Item:"Mobil"},
		&models.Order{IDCust : 1, Item:"Mobil"},
		}, nil)
		Order := NewOrderUsecase(Ordermocks)
		c := Order.GetAllOrderById(1)
		if len(c) != 10{		
			t.Errorf("Error harus 10, yang di dapat %d",len(c))
		}
	})
}