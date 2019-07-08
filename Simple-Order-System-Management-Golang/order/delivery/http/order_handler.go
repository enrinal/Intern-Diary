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
	e.GET("/orders", handler.FetchAllOrder)
	e.GET("/orders/:id", handler.FetchAllOrderByID)
	e.GET("/order/:id", handler.FetchOrderByID)
	e.GET("/checklimitorder/:id", handler.ChekcLimitOrder)
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
	idorder, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	id := int64(idorder)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listorders, err := order.OrderUsecase.GetAllOrderById(ctx, id)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, listorders)
}

func (order *OrderHandler) FetchOrderByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	if ctx != nil {
		ctx = context.Background()
	}
	orderid, err := order.OrderUsecase.GetOrderById(ctx, int64(id))
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, orderid)
}

func (order *OrderHandler) ChekcLimitOrder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	if ctx != nil {
		ctx = context.Background()
	}
	checklimit, err := order.OrderUsecase.CheckLimitOrder(ctx, int64(id))
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, checklimit)
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
