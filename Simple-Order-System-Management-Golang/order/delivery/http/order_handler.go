package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/order"
)

type ResponseError struct {
	Message string `json:"message"`
}

type OrderHandler struct {
	OrderUsecase order.Usecase
}

func NewOrderHandler(e *echo.Echo, order order.Usecase) {
	handler := &OrderHandler{
		OrderUsecase: order,
	}
	e.GET("/order", handler.FetchAllOrder)
	e.GET("/order", handler.FetchAllOrderByID)
}

func (order *OrderHandler) FetchAllOrder(c echo.Context) error {
	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listorders, err := order.OrderUsecase.GetAllOrder(ctx, int64(num))

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, listorders)
}

func (order *OrderHandler) FetchAllOrderByID(c echo.Context) error {
	ids := c.QueryParam("id")
	id, _ := strconv.Atoi(ids)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listorders, err := order.OrderUsecase.GetAllOrderById(ctx, int64(id))

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, listorders)
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
