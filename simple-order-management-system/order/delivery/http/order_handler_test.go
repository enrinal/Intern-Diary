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
	req, err := http.NewRequest(echo.GET, "/customer/order/"+strconv.Itoa(id), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("customer/order/:id")
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

func TestOrderCountId(t *testing.T) {
	var ordercount int
	var mockOrder models.Order
	err := faker.FakeData(&mockOrder)
	assert.NoError(t, err)

	mockUCase := new(mocks.Usecase)
	id := int(mockOrder.ID)

	mockUCase.On("CountOrderCust", mock.Anything, int64(id)).Return(ordercount, nil)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/countorder/"+strconv.Itoa(id), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("countorder/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))
	handler := OrderHandler{
		OrderUsecase: mockUCase,
	}
	err = handler.OrderCustCount(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)

}

func TestChangeOrderProcess(t *testing.T) {
	var mockOrder models.Order
	err := faker.FakeData(&mockOrder)
	assert.NoError(t, err)

	mockUCase := new(mocks.Usecase)
	id := int(mockOrder.ID)

	mockUCase.On("ChangeOrderProcess", mock.Anything, int64(id)).Return(nil)
	mockUCase.On("GetOrderById", mock.Anything, int64(id)).Return(&mockOrder, nil)
	e := echo.New()
	req, err := http.NewRequest(echo.PUT, "/changeorderprocess/"+strconv.Itoa(id), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("changeorderprocess/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))
	handler := OrderHandler{
		OrderUsecase: mockUCase,
	}
	err = handler.ChangeOrderProcess(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestChangeOrderSend(t *testing.T) {
	var mockOrder models.Order
	err := faker.FakeData(&mockOrder)
	assert.NoError(t, err)

	mockUCase := new(mocks.Usecase)
	id := int(mockOrder.ID)

	mockUCase.On("ChangeOrderSend", mock.Anything, int64(id)).Return(nil)
	mockUCase.On("GetOrderById", mock.Anything, int64(id)).Return(&mockOrder, nil)
	e := echo.New()
	req, err := http.NewRequest(echo.PUT, "/changeordersend/"+strconv.Itoa(id), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("changeordersend/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))
	handler := OrderHandler{
		OrderUsecase: mockUCase,
	}
	err = handler.ChangeOrderSend(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestChangeOrderDelivered(t *testing.T) {
	var mockOrder models.Order
	err := faker.FakeData(&mockOrder)
	assert.NoError(t, err)

	mockUCase := new(mocks.Usecase)
	id := int(mockOrder.ID)

	mockUCase.On("ChangeOrderDelivered", mock.Anything, int64(id)).Return(nil)
	mockUCase.On("GetOrderById", mock.Anything, int64(id)).Return(&mockOrder, nil)
	e := echo.New()
	req, err := http.NewRequest(echo.PUT, "/changeorderdelivered/"+strconv.Itoa(id), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("changeorderdelivered/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))
	handler := OrderHandler{
		OrderUsecase: mockUCase,
	}
	err = handler.ChangeOrderDelivered(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}
