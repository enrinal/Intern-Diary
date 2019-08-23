package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"

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
	req, err := http.NewRequest(echo.GET, "/api/v1/customers?num=1", strings.NewReader(""))
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

func TestFetchCustID(t *testing.T) {
	var mockCustomer models.Customer
	err := faker.FakeData(&mockCustomer)
	assert.NoError(t, err)

	mockUCase := new(mocks.Usecase)
	id := int(mockCustomer.ID)

	mockUCase.On("GetCustomerByID", mock.Anything, int64(id)).Return(&mockCustomer, nil)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/api/v1/customers/"+strconv.Itoa(id), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("customer/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))
	handler := CustomerHandler{
		CustUsecase: mockUCase,
	}
	err = handler.FetchCustomerByID(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)

}

func TestAdd(t *testing.T) {
	mockCustomer := models.Customer{
		ID:     int64(1),
		Name:   "Muhamad Enrinal",
		Status: 1,
	}

	tempMockCustomer := mockCustomer
	tempMockCustomer.ID = 0
	mockUcase := new(mocks.Usecase)

	j, err := json.Marshal(tempMockCustomer)
	assert.NoError(t, err)
	mockUcase.On("Add", mock.Anything, mock.AnythingOfType("*models.Customer")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/api/v1/customers", strings.NewReader(string(j)))
	assert.NoError(t, err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/customers")

	handler := CustomerHandler{
		CustUsecase: mockUcase,
	}
	err = handler.AddCustomer(c)

	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUcase.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	mockCustomer := models.Customer{
		ID:     int64(1),
		Name:   "Muhamad Enrinal",
		Status: 1,
	}

	tempMockCustomer := mockCustomer
	tempMockCustomer.ID = 0
	mockUcase := new(mocks.Usecase)

	j, err := json.Marshal(tempMockCustomer)
	assert.NoError(t, err)
	mockUcase.On("Update", mock.Anything, mock.AnythingOfType("*models.Customer")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.PUT, "/api/v1/customers", strings.NewReader(string(j)))
	assert.NoError(t, err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/customers")

	handler := CustomerHandler{
		CustUsecase: mockUcase,
	}
	err = handler.UpdateCustomer(c)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUcase.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	var mockCustomer models.Customer
	err := faker.FakeData(&mockCustomer)
	assert.NoError(t, err)

	mockUCase := new(mocks.Usecase)

	id := int(mockCustomer.ID)
	mockUCase.On("Delete", mock.Anything, int64(id)).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/api/v1/customers/"+strconv.Itoa(id), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("customers/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))
	handler := CustomerHandler{
		CustUsecase: mockUCase,
	}
	err = handler.DeleteCustomer(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockUCase.AssertExpectations(t)

}
