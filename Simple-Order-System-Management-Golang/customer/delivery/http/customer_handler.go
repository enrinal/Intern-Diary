package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/customer"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

type ResponseError struct {
	Message string `json:"message"`
}

type CustomerHandler struct {
	CustUsecase customer.Usecase
}

func NewCustomerHandler(e *echo.Echo, cust customer.Usecase) {
	handler := &CustomerHandler{
		CustUsecase: cust,
	}
	e.GET("/customers", handler.FetchCustomer)
	e.GET("/customers/:id", handler.FetchCustomerByID)
}

func (cust *CustomerHandler) FetchCustomer(c echo.Context) error {
	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listcusts, err := cust.CustUsecase.GetAllCustomer(ctx, int64(num))

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, listcusts)
}

func (cust *CustomerHandler) FetchCustomerByID(c echo.Context) error {
	idCust, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	id := int64(idCust)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	customer, err := cust.CustUsecase.GetCustomerByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, customer)

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
