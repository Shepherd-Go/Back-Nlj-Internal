package services

import (
	"context"
	"net/http"

	"github.com/BBCompanyca/Back-Nlj-Internal.git/db/models"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/dto"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/entity"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/repository"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/utils"
	"github.com/labstack/echo/v4"
)

type Employee interface {
	CreateEmployee(ctx context.Context, empl dto.EmployeeRequest) error
}

type employee struct {
	repoEmployee repository.Employee
	managePass   utils.Password
	logsError    utils.LogsError
}

func NewServiceEmployee(repoEmployee repository.Employee, managePass utils.Password, logsError utils.LogsError) Employee {
	return &employee{repoEmployee, managePass, logsError}
}

func (e *employee) CreateEmployee(ctx context.Context, empl dto.EmployeeRequest) error {

	passTemporary := e.managePass.GenerateTemporaryPassword()
	e.managePass.HashPassword(&passTemporary)

	empl.Password = passTemporary
	empl.Username = createUsername(empl.FirstName, empl.LastName, empl.Phone)
	parsePermissions(&empl.Permissions)

	buildEmployee := models.Employee{}
	buildEmployee.BuildCreateEmployeeModel(empl)

	if err := e.repoEmployee.CreateEmployee(ctx, buildEmployee); err != nil {
		e.logsError.InsertLogsError(err)
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "An unexpected error has occurred on the server"})
	}

	return nil
}

func createUsername(firstName, lastName, phone string) string {
	return firstName[:3] + lastName[:3] + "-" + phone[10:]
}

func parsePermissions(permissions *string) {
	switch *permissions {
	case "administrador":
		*permissions = "1"

	case "vendedor":
		*permissions = "2"
	}

}
