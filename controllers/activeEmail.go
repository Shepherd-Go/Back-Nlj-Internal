package controllers

import (
	"net/http"

	"github.com/Shepherd-Go/Back-Nlj-Internal.git/dtos"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/entity"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/services"
	"github.com/labstack/echo/v4"
)

type ActivateEmail interface {
	ActivateEmail(c echo.Context) error
}

type activateemail struct {
	svc services.ActivateEmail
}

func NewActiveEmailControler(svc services.ActivateEmail) ActivateEmail {
	return &activateemail{svc}
}

func (a *activateemail) ActivateEmail(c echo.Context) error {

	ctx := c.Request().Context()

	pass := dtos.ActivateEmail{}

	if err := c.Bind(&pass); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{Message: err.Error()})
	}

	if err := pass.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{Message: err.Error()})
	}

	empl, err := a.svc.ActivateEmail(ctx, pass)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, entity.Response{Message: "data employee session", Data: empl})
}
