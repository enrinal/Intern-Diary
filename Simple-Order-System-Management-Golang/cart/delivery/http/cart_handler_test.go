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
