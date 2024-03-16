package controller

import (
	"net/http"

	"github.com/BBCompanyca/Back-Nlj-Internal.git/dto"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/entity"
	"github.com/labstack/echo/v4"
)

type Employee interface {
	CreateEmployee(c echo.Context) error
}

type employee struct{}

func NewEmployeeController() Employee {
	return &employee{}
}

func (e *employee) CreateEmployee(c echo.Context) error {

	employee := dto.EmployeeRequest{}

	if err := c.Bind(&employee); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{Message: err.Error()})
	}

	if err := employee.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, entity.Response{Message: "employee created successfully.!!"})
}
