package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	customers "gitlab.warungpintar.co/enrinal/intern-diary/simple-order/customer/mocks"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
	orders "gitlab.warungpintar.co/enrinal/intern-diary/simple-order/order/mocks"
)

func TestGetAllOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		Ordermocks := new(orders.Repository)
		Ordermocks.On("GetAllOrder", mock.Anything, mock.AnythingOfType("int64")).Return([]*models.Order{
			&models.Order{ID: 1, Item: "Mobil"},
			&models.Order{ID: 1, Item: "Mobil"},
			&models.Order{ID: 1, Item: "Mobil"},
			&models.Order{ID: 1, Item: "Mobil"},
			&models.Order{ID: 1, Item: "Mobil"},
			&models.Order{ID: 1, Item: "Mobil"},
			&models.Order{ID: 1, Item: "Mobil"},
			&models.Order{ID: 1, Item: "Mobil"},
			&models.Order{ID: 1, Item: "Mobil"},
			&models.Order{ID: 1, Item: "Mobil"},
		}, nil)
		Order := NewOrderUsecase(Ordermocks, nil, time.Second*2)
		c, err := Order.GetAllOrder(context.TODO(), int64(10))
		assert.NoError(t, err)
		assert.Len(t, c, 10)
	})
}

func TestGetOrderById(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		Ordermocks := new(orders.Repository)
		Ordermocks.On("GetOrderById", mock.Anything, int64(2)).Return(&models.Order{ID: 2, Item: "Mobil", Status: 1}, nil)
		Order := NewOrderUsecase(Ordermocks, nil, time.Second*2)
		c, err := Order.GetOrderById(context.TODO(), int64(2))
		assert.NoError(t, err)
		assert.NotNil(t, c)
	})
}

func TestOrderProcess(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		Ordermocks := new(orders.Repository)
		Ordermocks.On("ChangeOrderProcess", mock.Anything, int64(2)).Once().Return(nil)
		Order := NewOrderUsecase(Ordermocks, nil, time.Second*2)
		err := Order.ChangeOrderProcess(context.TODO(), int64(2))
		assert.NoError(t, err)
	})
}

func TestOrderSend(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		Ordermocks := new(orders.Repository)
		Ordermocks.On("ChangeOrderSend", mock.Anything, int64(2)).Once().Return(nil)
		Order := NewOrderUsecase(Ordermocks, nil, time.Second*2)
		err := Order.ChangeOrderSend(context.TODO(), int64(2))
		assert.NoError(t, err)
	})
}

func TestOrderDelivered(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		Ordermocks :=new(orders.Repository)
		Ordermocks.On("ChangeOrderDelivered", mock.Anything, int64(2)).Once().Return(nil)
		Order := NewOrderUsecase(Ordermocks, nil, time.Second*2)
		err := Order.ChangeOrderDelivered(context.TODO(), int64(2))
		assert.NoError(t, err)
	})
}

func TestAllOrderByCustomerId(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		Ordermocks :=new(orders.Repository)
		Ordermocks.On("GetAllOrderById", mock.Anything, int64(1)).Return([]*models.Order{
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
		}, nil)
		Order := NewOrderUsecase(Ordermocks, nil, time.Second*2)
		c, err := Order.GetAllOrderById(context.TODO(), int64(1))
		assert.NoError(t, err)
		assert.Len(t, c, 10)
	})
}

func TestCountAllOrderByCustomerId(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		Ordermocks := new(orders.Repository)
		Ordermocks.On("GetAllOrderById", mock.Anything, int64(1)).Return([]*models.Order{
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
		}, nil)
		Order := NewOrderUsecase(Ordermocks, nil, time.Second*2)
		c, err := Order.CountOrderCust(context.TODO(), int64(1))
		assert.NoError(t, err)
		assert.Equal(t, c, 10)
	})
}

func TestCheckLimitOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		Ordermocks := new(orders.Repository)
		Ordermocks.On("GetAllOrderById", mock.Anything, int64(1)).Return([]*models.Order{
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
		}, nil)
		customermocks := new(customers.Repository)
		customermocks.On("GetCustomerById", mock.Anything, int64(1)).Return(&models.Customer{ID: 1, Name: "Name", Status: 2}, nil)
		Order := NewOrderUsecase(Ordermocks, customermocks, time.Second*2)
		c, err := Order.CheckLimitOrder(context.TODO(), int64(1))
		assert.NoError(t, err)
		assert.Equal(t, c, true)
	})
	t.Run("success", func(t *testing.T) {
		Ordermocks := new(orders.Repository)
		Ordermocks.On("GetAllOrderById", mock.Anything, int64(1)).Return([]*models.Order{
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
			&models.Order{IDCust: 1, Item: "Mobil"},
		}, nil)
		customermocks := new(customers.Repository)
		customermocks.On("GetCustomerById", mock.Anything, int64(1)).Return(&models.Customer{ID: 1, Name: "Name", Status: 1}, nil)
		Order := NewOrderUsecase(Ordermocks, customermocks, time.Second*2)
		c, err := Order.CheckLimitOrder(context.TODO(), int64(1))
		assert.NoError(t, err)
		assert.Equal(t, c, false)
	})
}
