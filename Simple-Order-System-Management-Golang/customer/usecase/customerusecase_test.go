package usecase

import (
	"testing"

	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/customer/mocks"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

func TestGetAllCustomer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		customermocks := &mocks.Usecase{}
		customermocks.On("Fetch").Return([]*models.Customer{
			&models.Customer{ID: 1, Name: "Muhamad Enrinal"},
			&models.Customer{ID: 1, Name: "Muhamad Enrinal"},
			&models.Customer{ID: 1, Name: "Muhamad Enrinal"},
			&models.Customer{ID: 1, Name: "Muhamad Enrinal"},
			&models.Customer{ID: 1, Name: "Muhamad Enrinal"},
			&models.Customer{ID: 1, Name: "Muhamad Enrinal"},
			&models.Customer{ID: 1, Name: "Muhamad Enrinal"},
			&models.Customer{ID: 1, Name: "Muhamad Enrinal"},
			&models.Customer{ID: 1, Name: "Muhamad Enrinal"},
			&models.Customer{ID: 1, Name: "Muhamad Enrinal"},
		}, nil)
		customer := NewCustomerUsecase(customermocks)
		c := customer.GetAllCustomer()
		if len(c) != 10 {
			t.Errorf("Error harus 10, yang di dapat %d", len(c))
		}
	})
}

func TestGetCustomerById(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		customermocks := &mocks.Usecase{}
		customermocks.On("GetById", int64(2)).Return(&models.Customer{ID: 2, Name: "Muhamad Enrinal", Status: 1}, nil)
		customer := NewCustomerUsecase(customermocks)
		c := customer.GetCustomerById(int64(2))

		if c == nil {
			t.Errorf("Error not found id")
		}
	})
}
