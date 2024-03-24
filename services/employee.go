package services

import (
	"context"
	"errors"
	"net/http"

	"github.com/BBCompanyca/Back-Nlj-Internal.git/db/models"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/db/repository"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/dtos"

	"github.com/BBCompanyca/Back-Nlj-Internal.git/entity"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/utils"
	"github.com/labstack/echo/v4"
)

type Employee interface {
	CreateEmployee(ctx context.Context, empl dtos.RegisterEmployee) error
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

	emplModel, err := e.repoEmployee.SearchEmployeeByEmail(ctx, empl.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "An unexpected error has occurred on the server"})
	}

	if emplModel.ID != "" {
		return echo.NewHTTPError(http.StatusConflict, entity.Response{Message: "email address not available, a user already occupies it"})
	}

	passTemporary := e.managePass.GenerateTemporaryPassword()
	e.managePass.HashPassword(&passTemporary)

	empl.Password = passTemporary
	if err = parsePermissions(&empl.Permissions); err != nil {
		return echo.NewHTTPError(http.StatusConflict, entity.Response{Message: err.Error()})
	}

	buildEmployee := models.Employee{}
	buildEmployee.BuildCreateEmployeeModel(empl)

	if err := e.repoEmployee.CreateEmployee(ctx, buildEmployee); err != nil {
		e.logsError.InsertLogsError(err)
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "An unexpected error has occurred on the server"})
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
