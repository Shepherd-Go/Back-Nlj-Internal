package services

import (
	"context"
	"log"
	"net/http"

	"github.com/Shepherd-Go/Back-Nlj-Internal.git/db/models"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/db/repository"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/dtos"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/entity"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PaymentMethods interface {
	RegisterMobilePayment(ctx context.Context, paymobl dtos.PaymentMobile) error
	SearchAllMobilePayment(ctx context.Context, idEmployee uuid.UUID) (dtos.PaymentMobileResponse, error)
	DeleteMobilePayment(ctx context.Context, id uuid.UUID) error
}

type paymentmethods struct {
	repoEmployee       repository.Employee
	repoPaymentMethods repository.PaymentMethods
}

func NewPaymentMethodsService(repoEmployee repository.Employee, repoPaymentMethods repository.PaymentMethods) PaymentMethods {
	return &paymentmethods{
		repoEmployee,
		repoPaymentMethods}
}

func (pmthds *paymentmethods) RegisterMobilePayment(ctx context.Context, paymobl dtos.PaymentMobile) error {

	id := ctx.Value("id").(string)

	empl, err := pmthds.repoEmployee.SearchEmployeeByID(ctx, uuid.MustParse(id))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	if empl.ID == uuid.Nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{Message: "there is no employee with this id"})
	}

	payMobil, err := pmthds.repoPaymentMethods.SearchPaymentMobileByEmployeeID(ctx, empl.ID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	if payMobil.ID != uuid.Nil {
		return echo.NewHTTPError(http.StatusConflict, entity.Response{Message: "this employee already has this payment method registered in the system."})
	}

	pm, err := pmthds.repoPaymentMethods.SearchPaymentMobileByID(ctx, paymobl.ID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	if pm.ID != uuid.Nil {
		return echo.NewHTTPError(http.StatusConflict, entity.Response{Message: "there is already a mobile payment with this id"})
	}

	paymobl.EmployeeID = uuid.MustParse(id)
	buildModel := models.PaymentMobile{}
	buildModel.BuildRegisterPaymentMobile(paymobl)

	if err := pmthds.repoPaymentMethods.RegisterPaymentMobile(ctx, buildModel); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	return nil
}

func (pmthds *paymentmethods) SearchAllMobilePayment(ctx context.Context, idEmployee uuid.UUID) (dtos.PaymentMobileResponse, error) {

	pmethods, err := pmthds.repoPaymentMethods.SearchPaymentMobileByEmployeeID(ctx, idEmployee)
	if err != nil {
		log.Println(err)
		return dtos.PaymentMobileResponse{}, echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	return pmethods, nil
}

func (pmthds *paymentmethods) DeleteMobilePayment(ctx context.Context, id uuid.UUID) error {

	idEmployee := ctx.Value("id").(string)

	payMobile, err := pmthds.SearchAllMobilePayment(ctx, uuid.MustParse(idEmployee))
	if err != nil {
		return err
	}

	if payMobile.ID == uuid.Nil {
		return echo.NewHTTPError(http.StatusNotFound, entity.Response{Message: "this employee does not have mobile payment registered in the system"})
	}

	if payMobile.ID != id {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{Message: "this mobile payment does not belong to this user."})
	}

	if err := pmthds.repoPaymentMethods.DeletePaymentMobile(ctx, id); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	return nil
}
