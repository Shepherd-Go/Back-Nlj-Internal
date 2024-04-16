package services

import (
	"context"
	"net/http"

	"github.com/Shepherd-Go/Back-Nlj-Internal.git/db/repository"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/dtos"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/entity"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Session interface {
	Session(ctx context.Context, login dtos.LogIn) (dtos.Session, error)
}

type session struct {
	logIn repository.LogIn
	pass  utils.Password
}

func NewSessionService(logIn repository.LogIn, pass utils.Password) Session {
	return &session{logIn, pass}
}

func (s *session) Session(ctx context.Context, login dtos.LogIn) (dtos.Session, error) {

	empl, err := s.logIn.SearchEmployeByEmailOrUsername(ctx, login.Identifier)
	if err != nil {
		return dtos.Session{}, echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	if empl.ID == uuid.Nil {
		return dtos.Session{}, echo.NewHTTPError(http.StatusNotFound, entity.Response{Message: "incorrect access data"})
	}

	if !(s.pass.CheckPasswordHash(empl.Password, login.Password)) {
		return dtos.Session{}, echo.NewHTTPError(http.StatusNotFound, entity.Response{Message: "incorrect access data"})
	}

	if !*empl.Status {
		return dtos.Session{}, echo.NewHTTPError(http.StatusForbidden, entity.Response{Message: "inactive employee"})
	}

	empl.Password = nil
	empl.Status = nil

	return empl, nil
}
