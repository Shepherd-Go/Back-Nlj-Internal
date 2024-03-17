package controllers

import (
	"net/http"

	"github.com/BBCompanyca/Back-Nlj-Internal.git/dto"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/entity"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/services"
	"github.com/labstack/echo/v4"
)

type Employee interface {
	CreateEmployee(c echo.Context) error
}

type employee struct {
	emplService services.Employee
}

func NewEmployeeController(emplService services.Employee) Employee {
	return &employee{emplService}
}

func (e *employee) CreateEmployee(c echo.Context) error {

	ctx := c.Request().Context()

	empl := dto.EmployeeRequest{}

	if err := c.Bind(&empl); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{Message: err.Error()})
	}

	if err := empl.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{Message: err.Error()})
	}

	if err := e.emplService.CreateEmployee(ctx, empl); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, entity.Response{Message: "employee created successfully.!!"})
}
