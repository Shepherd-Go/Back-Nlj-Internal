package controllers

import (
	"net/http"

	"github.com/Shepherd-Go/Back-Nlj-Internal.git/dtos"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/entity"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PaymentMethods interface {
	RegisterMobilePayment(c echo.Context) error
	SearchAllMobilePayment(c echo.Context) error
}

type paymentmethods struct {
	servPaymentMethods services.PaymentMethods
}

func NewPaymentMethodsController(servPaymentMethods services.PaymentMethods) PaymentMethods {
	return &paymentmethods{servPaymentMethods}
}

func (pmthds *paymentmethods) RegisterMobilePayment(c echo.Context) error {

	ctx := c.Request().Context()

	payMobile := dtos.PaymentMobile{}

	if err := c.Bind(&payMobile); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{Message: err.Error()})
	}

	if err := payMobile.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{Message: err.Error()})
	}

	if err := pmthds.servPaymentMethods.RegisterMobilePayment(ctx, payMobile); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, entity.Response{Message: "payment mobile registered successfully.!!"})
}

func (pmthds *paymentmethods) SearchAllMobilePayment(c echo.Context) error {

	ctx := c.Request().Context()

	id := c.Param("id")

	idUUID, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, entity.Response{Message: "format id invalid"})
	}

	pmethods, err := pmthds.servPaymentMethods.SearchAllMobilePayment(ctx, idUUID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, entity.Response{Message: "all mobile payments", Data: pmethods})

}
