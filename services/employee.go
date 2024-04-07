package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/BBCompanyca/Back-Nlj-Internal.git/db/models"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/db/repository"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/dtos"
	"github.com/google/uuid"

	"github.com/BBCompanyca/Back-Nlj-Internal.git/entity"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/utils"
	"github.com/labstack/echo/v4"
)

type Employee interface {
	CreateEmployee(ctx context.Context, empl dtos.RegisterEmployee) error
	GetEmployees(ctx context.Context) (dtos.Employees, error)
	UpdateEmployees(ctx context.Context, id uuid.UUID, empl dtos.UpdateEmployee) error
	DeleteEmployee(ctx context.Context, id uuid.UUID) error
}

type employee struct {
	repoEmployee repository.Employee
	managePass   utils.Password
	logsError    utils.LogsError
}

func NewServiceEmployee(repoEmployee repository.Employee, managePass utils.Password, logsError utils.LogsError) Employee {
	return &employee{repoEmployee, managePass, logsError}
}

func (e *employee) CreateEmployee(ctx context.Context, empl dtos.RegisterEmployee) error {

	/*id := ctx.Value("id").(string)
	permissions := ctx.Value("permissions").(string)

	if permissions != "administrator" {
		return echo.NewHTTPError(http.StatusUnauthorized, entity.Response{Message: "unauthorized"})
	}*/

	emplModel, err := e.repoEmployee.SearchEmployeeByEmail(ctx, empl.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	if emplModel.ID != uuid.Nil {
		return echo.NewHTTPError(http.StatusConflict, entity.Response{Message: "email address not available, a user already occupies it"})
	}

	passTemporary := e.managePass.GenerateTemporaryPassword()
	fmt.Println(passTemporary)
	e.managePass.HashPassword(&passTemporary)
	empl.Password = passTemporary

	if err = parsePermissions(&empl.Permissions); err != nil {
		return echo.NewHTTPError(http.StatusConflict, entity.Response{Message: err.Error()})
	}

	//empl.Created_by = id
	//empl.Updated_by = id

	buildEmployee := models.Employee{}
	buildEmployee.BuildCreateEmployeeModel(empl)

	if err := e.repoEmployee.CreateEmployee(ctx, buildEmployee); err != nil {
		e.logsError.InsertLogsError(err)
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	return nil
}

func (e *employee) GetEmployees(ctx context.Context) (dtos.Employees, error) {

	/*permissions := ctx.Value("permissions")

	if permissions != "administrator" {
		return dtos.Employees{}, echo.NewHTTPError(http.StatusUnauthorized, entity.Response{Message: "unauthorized"})
	}*/

	empls, err := e.repoEmployee.SearchAllEmployees(ctx)
	if err != nil {
		return dtos.Employees{}, echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	return empls, nil
}

func (e *employee) UpdateEmployees(ctx context.Context, id uuid.UUID, empl dtos.UpdateEmployee) error {

	/*idToken := ctx.Value("id").(string)
	permissions := ctx.Value("permissions")

	if permissions != "administrator" {
		return echo.NewHTTPError(http.StatusUnauthorized, entity.Response{Message: "unauthorized"})
	}*/

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
	//empl.Updated_by = idToken
	buildModelEmploye := models.Employee{}
	buildModelEmploye.BuildUpdatedEmployeeModel(empl, id)

	if err := e.repoEmployee.UpdateEmployee(ctx, buildModelEmploye, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	return nil
}

func (e *employee) DeleteEmployee(ctx context.Context, id uuid.UUID) error {

	emplModel, err := e.repoEmployee.SearchEmployeeByID(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	if emplModel.ID == uuid.Nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{Message: "there is no employee with this id"})
	}

	if err := e.repoEmployee.DeleteEmployee(ctx, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	return nil
}

func parsePermissions(permissions *string) error {
	switch *permissions {
	case "administrator":
		*permissions = "1"
	case "seller":
		*permissions = "2"
	default:
		return errors.New("send permissions do not exist")
	}

	return nil
}
