package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/cart"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
	validator "gopkg.in/go-playground/validator.v9"
)

type ResponseError struct {
	Message string `json:"message"`
}

type CartHandler struct {
	CartUsecase cart.Usecase
}

func NewCartHandler(e *echo.Echo, cart cart.Usecase) {
	handler := &CartHandler{
		CartUsecase: cart,
	}
	e.GET("/cart/:id", handler.FetchCustCart)
	e.POST("/cart", handler.Store)
	e.DELETE("/cart", handler.Delete)
}

func (cart *CartHandler) FetchCustCart(c echo.Context) error {
	idcust, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	id := int64(idcust)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listcart, err := cart.CartUsecase.GetCustCart(id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, listcart)
}

func (a *CartHandler) Store(c echo.Context) error {
	var cart models.Cart
	err := c.Bind(&cart)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&cart); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = a.CartUsecase.AddItem(cart)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, cart)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func isRequestValid(m *models.Cart) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *CartHandler) Delete(c echo.Context) error {
	var cart models.Cart
	err := c.Bind(&cart)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&cart); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = a.CartUsecase.RemoveItem(cart)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, cart)
}
