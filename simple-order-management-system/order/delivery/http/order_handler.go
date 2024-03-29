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
	e.GET("/api/v1/orders", handler.FetchAllOrder)
	e.GET("/api/v1/orders/customer/:id", handler.FetchAllOrderByID)
	e.GET("/api/v1/orders/:id", handler.FetchOrderByID)
	e.GET("/api/v1/orders/:id/limitorder", handler.ChekcLimitOrder)
	e.GET("/api/v1/orders/:id/countorder", handler.OrderCustCount)
	e.PUT("/api/v1/orders/:id/process", handler.ChangeOrderProcess)
	e.PUT("/api/v1/orders/:id/send", handler.ChangeOrderSend)
	e.PUT("/api/v1/orders/:id/delivered", handler.ChangeOrderDelivered)
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

func (order *OrderHandler) OrderCustCount(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	if ctx != nil {
		ctx = context.Background()
	}
	countorder, err := order.OrderUsecase.CountOrderCust(ctx, int64(id))
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, countorder)
}

func (order *OrderHandler) ChangeOrderProcess(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	if ctx != nil {
		ctx = context.Background()
	}
	err = order.OrderUsecase.ChangeOrderProcess(ctx, int64(id))
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	orderid, err := order.OrderUsecase.GetOrderById(ctx, int64(id))
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, orderid)
}

func (order *OrderHandler) ChangeOrderSend(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	if ctx != nil {
		ctx = context.Background()
	}
	err = order.OrderUsecase.ChangeOrderSend(ctx, int64(id))
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	orderid, err := order.OrderUsecase.GetOrderById(ctx, int64(id))
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, orderid)
}

func (order *OrderHandler) ChangeOrderDelivered(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	if ctx != nil {
		ctx = context.Background()
	}
	err = order.OrderUsecase.ChangeOrderDelivered(ctx, int64(id))
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	orderid, err := order.OrderUsecase.GetOrderById(ctx, int64(id))
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, orderid)
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
