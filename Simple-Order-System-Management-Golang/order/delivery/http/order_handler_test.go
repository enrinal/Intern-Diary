package http

import (
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
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/order/mocks"
)

func TestFetch(t *testing.T) {
	var mockOrder models.Order
	err := faker.FakeData(&mockOrder)
	assert.NoError(t, err)
	mockUCase := new(mocks.Usecase)
	mockListOrder := make([]*models.Order, 0)
	mockListOrder = append(mockListOrder, &mockOrder)
	num := 1
	mockUCase.On("GetAllOrder", mock.Anything, int64(num)).Return(mockListOrder, nil)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/orders?num=1", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := OrderHandler{
		OrderUsecase: mockUCase,
	}
	err = handler.FetchAllOrder(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestFetchAllOrderByID(t *testing.T) {
	var mockOrder models.Order
	err := faker.FakeData(&mockOrder)
	assert.NoError(t, err)

	mockUCase := new(mocks.Usecase)
	id := int(mockOrder.ID)
	mockListOrder := make([]*models.Order, 0)
	mockListOrder = append(mockListOrder, &mockOrder)
	mockUCase.On("GetAllOrderById", mock.Anything, int64(id)).Return(mockListOrder, nil)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/orders/"+strconv.Itoa(id), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("orders/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))
	handler := OrderHandler{
		OrderUsecase: mockUCase,
	}
	err = handler.FetchAllOrderByID(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestFetchOrderID(t *testing.T) {
	var mockOrder models.Order
	err := faker.FakeData(&mockOrder)
	assert.NoError(t, err)

	mockUCase := new(mocks.Usecase)
	id := int(mockOrder.ID)

	mockUCase.On("GetOrderById", mock.Anything, int64(id)).Return(&mockOrder, nil)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/order/"+strconv.Itoa(id), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("order/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))
	handler := OrderHandler{
		OrderUsecase: mockUCase,
	}
	err = handler.FetchOrderByID(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)

}

func TestCheckLimitOrder(t *testing.T) {
	var limitorder bool
	var mockOrder models.Order
	err := faker.FakeData(&mockOrder)
	assert.NoError(t, err)

	mockUCase := new(mocks.Usecase)
	id := int(mockOrder.ID)

	mockUCase.On("CheckLimitOrder", mock.Anything, int64(id)).Return(limitorder, nil)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/checklimitorder/"+strconv.Itoa(id), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("checklimitorder/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))
	handler := OrderHandler{
		OrderUsecase: mockUCase,
	}
	err = handler.ChekcLimitOrder(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)

}
