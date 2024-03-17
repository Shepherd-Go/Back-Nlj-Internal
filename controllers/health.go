package controllers

import (
	"net/http"

	"github.com/BBCompanyca/Back-Nlj-Internal.git/entity"
	"github.com/labstack/echo/v4"
)

type Health interface {
	Health(c echo.Context) error
}

type health struct{}

func NewHealthController() Health {
	return &health{}
}

func (h *health) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, entity.Response{Message: "service running successfully..."})
}
