package http

import (
	"net/http"
	"net/http/httptest"

	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/customer/mocks"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

func TestFetch(t *testing.T) {
	var mockCustomer models.Customer
	err := faker.FakeData(&mockCustomer)
	assert.NoError(t, err)
	mockUCase := new(mocks.Usecase)
	mockListCustomer := make([]*models.Customer, 0)
	mockListCustomer = append(mockListCustomer, &mockCustomer)
	num := 1
	mockUCase.On("GetAllCustomer", mock.Anything, int64(num)).Return(mockListCustomer, nil)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/customers?num=1", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := CustomerHandler{
		CustUsecase: mockUCase,
	}
	err = handler.FetchCustomer(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestFetchError(t *testing.T) {
	mockUCase := new(mocks.Usecase)
	num := 1
	mockUCase.On("GetAllCustomer", mock.Anything, int64(num)).Return(nil, models.ErrInternalServerError)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/customers?num=1", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := CustomerHandler{
		CustUsecase: mockUCase,
	}
	err = handler.FetchCustomer(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertExpectations(t)
}
