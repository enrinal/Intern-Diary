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
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/cart/mocks"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

func TestFetchCustCart(t *testing.T) {
	var mockCart models.Cart
	err := faker.FakeData(&mockCart)
	assert.NoError(t, err)

	mockUCase := new(mocks.Usecase)
	id := int(mockCart.IDCust)
	mockListCaart := make([]models.Cart, 0)
	mockListCaart = append(mockListCaart, mockCart)
	mockUCase.On("GetCustCart", int64(id)).Return(mockListCaart, nil)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/cart/"+strconv.Itoa(id), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("cart/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))
	handler := CartHandler{
		CartUsecase: mockUCase,
	}
	err = handler.FetchCustCart(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestStore(t *testing.T) {
	mockItem := make([]models.Item, 0)
	mockCart := models.Cart{
		IDCart: int64(1),
		IDCust: int64(2),
		Items:  mockItem,
		Qty:    int64(1),
		Prize:  int64(2),
	}

	tempmockCart := mockCart
	tempmockCart.IDCart = 1
	mockUCase := new(mocks.Usecase)

	j, err := json.Marshal(tempmockCart)
	assert.NoError(t, err)

	mockUCase.On("AddItem", mock.AnythingOfType("models.Cart")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/cart", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cart")

	handler := CartHandler{
		CartUsecase: mockUCase,
	}
	err = handler.Store(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockItem := make([]models.Item, 0)
	mockCart := models.Cart{
		IDCart: int64(1),
		IDCust: int64(2),
		Items:  mockItem,
		Qty:    int64(1),
		Prize:  int64(2),
	}

	tempmockCart := mockCart
	tempmockCart.IDCart = 1
	mockUCase := new(mocks.Usecase)

	j, err := json.Marshal(tempmockCart)
	assert.NoError(t, err)

	mockUCase.On("RemoveItem", mock.AnythingOfType("models.Cart")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/cart", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cart")

	handler := CartHandler{
		CartUsecase: mockUCase,
	}
	err = handler.Delete(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertExpectations(t)
}
