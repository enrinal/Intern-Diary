package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/customer/mocks"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

func TestGetAllCustomer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		customermocks := &mocks.Repository{}
		customermocks.On("Fetch", mock.Anything, mock.AnythingOfType("int64")).Return([]*models.Customer{
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
		customer := NewCustomerUsecase(customermocks, time.Second*2)
		c, err := customer.GetAllCustomer(context.TODO(), int64(10))
		assert.NoError(t, err)
		assert.Len(t, c, 10)
	})
}

func TestGetCustomerById(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		customermocks := &mocks.Repository{}
		customermocks.On("GetCustomerById", mock.Anything, mock.AnythingOfType("int64")).Return(&models.Customer{ID: 2, Name: "Muhamad Enrinal", Status: 1}, nil)
		customer := NewCustomerUsecase(customermocks, time.Second*2)
		c, err := customer.GetCustomerByID(context.TODO(), int64(2))

		assert.NoError(t, err)
		assert.NotNil(t, c)
	})
}
