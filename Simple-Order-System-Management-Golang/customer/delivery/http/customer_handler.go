package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/customer"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
	validator "gopkg.in/go-playground/validator.v9"
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
	e.GET("/api/v1/customers", handler.FetchCustomer)
	e.GET("/api/v1/customers/:id", handler.FetchCustomerByID)
	e.POST("/api/v1/customers", handler.AddCustomer)
	e.PUT("/api/v1/customers", handler.UpdateCustomer)
	e.DELETE("/api/v1/customers", handler.DeleteCustomer)
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

func (cust *CustomerHandler) AddCustomer(c echo.Context) error {
	var customer models.Customer
	err := c.Bind(&customer)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	if ok, err := isRequestValid(&customer); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	err = cust.CustUsecase.Add(ctx, &customer)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, customer)
}

func (cust *CustomerHandler) UpdateCustomer(c echo.Context) error {
	var customer models.Customer
	err := c.Bind(&customer)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	if ok, err := isRequestValid(&customer); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	err = cust.CustUsecase.Update(ctx, &customer)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, customer)
}

func (cust *CustomerHandler) DeleteCustomer(c echo.Context) error {
	idcust, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = cust.CustUsecase.Delete(ctx, int64(idcust))
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)

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

func isRequestValid(m *models.Customer) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
