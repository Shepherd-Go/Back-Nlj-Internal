package services

import (
	"context"
	"errors"
	"net/http"

	"github.com/Shepherd-Go/Back-Nlj-Internal.git/db/models"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/db/repository"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/dtos"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/entity"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Employee interface {
	CreateEmployee(ctx context.Context, empl dtos.RegisterEmployee) error
	GetEmployees(ctx context.Context) (dtos.Employees, error)
	UpdateEmployees(ctx context.Context, id uuid.UUID, empl dtos.UpdateEmployee) error
	DeleteEmployee(ctx context.Context, id uuid.UUID) error
	ForgotPassword(ctx context.Context, email string) error
}

type employee struct {
	repoEmployee repository.Employee
	managePass   utils.Password
	sendEmail    utils.SendEmail
}

func NewServiceEmployee(repoEmployee repository.Employee, managePass utils.Password, sendEmail utils.SendEmail) Employee {
	return &employee{repoEmployee, managePass, sendEmail}
}

func (e *employee) CreateEmployee(ctx context.Context, empl dtos.RegisterEmployee) error {

	//_ := ctx.Value("id").(string)

	emplModel, err := e.repoEmployee.SearchEmployeeByEmail(ctx, empl.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	if emplModel.ID != uuid.Nil {
		return echo.NewHTTPError(http.StatusConflict, entity.Response{Message: "email address not available, a user already occupies it"})
	}

	passTemporary := e.managePass.GenerateTemporaryPassword()
	empl.Password = passTemporary
	e.managePass.HashPassword(&empl.Password)

	if err = parsePermissions(&empl.Permissions); err != nil {
		return echo.NewHTTPError(http.StatusConflict, entity.Response{Message: err.Error()})
	}

	buildEmployee := models.Employee{}
	buildEmployee.BuildCreateEmployeeModel(empl)

	if err := e.repoEmployee.CreateEmployee(ctx, buildEmployee); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	go e.sendEmail.EmployeeRegistered(buildEmployee.Email, buildEmployee.FirstName, buildEmployee.Username, passTemporary)

	return nil
}

func (e *employee) GetEmployees(ctx context.Context) (dtos.Employees, error) {

	empls, err := e.repoEmployee.SearchAllEmployees(ctx)
	if err != nil {
		return dtos.Employees{}, echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	return empls, nil
}

func (e *employee) UpdateEmployees(ctx context.Context, id uuid.UUID, empl dtos.UpdateEmployee) error {

	emplModel, err := e.repoEmployee.SearchEmployeeByID(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	if emplModel.ID == uuid.Nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{Message: "there is no employee with this id"})
	}

	if emplModel.Email != empl.Email {
		auxEmplModel, err := e.repoEmployee.SearchEmployeeByEmailAndNotID(ctx, id, empl.Email)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
		}

		if auxEmplModel.ID != uuid.Nil {
			return echo.NewHTTPError(http.StatusConflict, entity.Response{Message: "email address not available, a user already occupies it"})
		}
	}

	if err = parsePermissions(&empl.Permissions); err != nil {
		return echo.NewHTTPError(http.StatusConflict, entity.Response{Message: err.Error()})
	}

	empl.ID = id
	buildModelEmploye := models.Employee{}
	buildModelEmploye.BuildUpdatedEmployeeModel(empl)

	if err := e.repoEmployee.UpdateEmployee(ctx, buildModelEmploye, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	return nil
}

func (e *employee) DeleteEmployee(ctx context.Context, id uuid.UUID) error {

	idToken := ctx.Value("id").(string)

	if idToken == id.String() {
		return echo.NewHTTPError(http.StatusConflict, entity.Response{Message: "you cannot delete your own account"})
	}

	emplModel, err := e.repoEmployee.SearchEmployeeByID(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	if emplModel.ID == uuid.Nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{Message: "there is no employee with this id"})
	}

	if err := e.repoEmployee.DeleteEmployee(ctx, id, idToken); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	return nil
}

func (e *employee) ForgotPassword(ctx context.Context, email string) error {

	empl, err := e.repoEmployee.SearchEmployeeByEmail(ctx, email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	if empl.ID == uuid.Nil {
		return echo.NewHTTPError(http.StatusNotFound, entity.Response{Message: "there is no account with this email"})
	}

	passTemporary := e.managePass.GenerateTemporaryPassword()
	passAxu := passTemporary
	e.managePass.HashPassword(&passAxu)

	if err := e.repoEmployee.ForgotPassword(ctx, empl.ID, passAxu); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	go e.sendEmail.ForgotPassword(empl.Email, passTemporary)

	return nil
}

func parsePermissions(permissions *string) error {
	switch *permissions {
	case "administrador":
		*permissions = "1"
	case "vendedor":
		*permissions = "2"
	default:
		return errors.New("send permissions do not exist")
	}

	return nil
}
