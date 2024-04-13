package controllers

import (
	"net/http"

	"github.com/Shepherd-Go/Back-Nlj-Internal.git/dtos"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/entity"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Employee interface {
	CreateEmployee(c echo.Context) error
	GetEmployees(c echo.Context) error
	UpdateEmployee(c echo.Context) error
	DeleteEmployee(c echo.Context) error
}

type employee struct {
	emplService services.Employee
}

func NewEmployeeController(emplService services.Employee) Employee {
	return &employee{emplService}
}

func (e *employee) CreateEmployee(c echo.Context) error {

	ctx := c.Request().Context()

	empl := dtos.RegisterEmployee{}

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

func (e *employee) GetEmployees(c echo.Context) error {

	ctx := c.Request().Context()

	empls, err := e.emplService.GetEmployees(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, entity.Response{Message: "all employees ok", Data: empls})
}

func (e *employee) UpdateEmployee(c echo.Context) error {

	ctx := c.Request().Context()

	id := c.QueryParam("id")

	idUUID, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, entity.Response{Message: "format id invalid"})
	}

	empl := dtos.UpdateEmployee{}

	if err := c.Bind(&empl); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{Message: err.Error()})
	}

	if err := empl.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{Message: err.Error()})
	}

	if err = e.emplService.UpdateEmployees(ctx, idUUID, empl); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, entity.Response{Message: "employee updated successfully.!!"})
}

func (e *employee) DeleteEmployee(c echo.Context) error {

	ctx := c.Request().Context()

	id := c.QueryParam("id")

	idUUID, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, entity.Response{Message: "format id invalid"})
	}

	if err := e.emplService.DeleteEmployee(ctx, idUUID); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, entity.Response{Message: "employee deleted successfully.!!"})
}
