package services

import (
	"context"
	"net/http"

	"github.com/BBCompanyca/Back-Nlj-Internal.git/db/repository"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/dtos"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/entity"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Session interface {
	Session(ctx context.Context, login dtos.Login) (dtos.EmployeeResponse, error)
}

type session struct {
	repoEmployee repository.Employee
	pass         utils.Password
}

func NewSessionService(repoEmployee repository.Employee, pass utils.Password) Session {
	return &session{repoEmployee, pass}
}

func (s *session) Session(ctx context.Context, login dtos.Login) (dtos.EmployeeResponse, error) {

	empl, err := s.repoEmployee.SearchEmployeByEmailOrUsername(ctx, login.Identifier)
	if err != nil {
		return dtos.EmployeeResponse{}, echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	if empl.ID == uuid.Nil {
		return dtos.EmployeeResponse{}, echo.NewHTTPError(http.StatusNotFound, entity.Response{Message: "incorrect access data"})
	}

	if !(s.pass.CheckPasswordHash(empl.Password, login.Password)) {
		return dtos.EmployeeResponse{}, echo.NewHTTPError(http.StatusNotFound, entity.Response{Message: "incorrect access data"})
	}

	if !*empl.Status {
		return dtos.EmployeeResponse{}, echo.NewHTTPError(http.StatusForbidden, entity.Response{Message: "inactive employee"})
	}

	empl.Password = nil

	return empl, nil
}
