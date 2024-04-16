package controllers

import (
	"net/http"

	"github.com/Shepherd-Go/Back-Nlj-Internal.git/dtos"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/entity"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/services"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/utils"
	"github.com/labstack/echo/v4"
)

type Session interface {
	Session(c echo.Context) error
}

type session struct {
	svcSession services.Session
	jwt        utils.JWT
}

func NewSessionController(svcSession services.Session, jwt utils.JWT) Session {
	return &session{svcSession, jwt}
}

func (s *session) Session(c echo.Context) error {

	ctx := c.Request().Context()

	login := dtos.LogIn{}

	if err := c.Bind(&login); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{Message: err.Error()})
	}

	if err := login.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{Message: err.Error()})
	}

	sessionEmployee, err := s.svcSession.Session(ctx, login)
	if err != nil {
		return err
	}

	token, err := s.jwt.SignedLoginToken(sessionEmployee)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{Message: "an unexpected error has occurred on the server"})
	}

	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    token,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Path:     "/",
	}

	c.SetCookie(cookie)

	return c.JSON(http.StatusCreated, entity.Response{Message: "session created successfully.!!", Data: sessionEmployee})
}
