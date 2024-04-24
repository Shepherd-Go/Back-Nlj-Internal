package middleware

import (
	"context"
	"net/http"

	"github.com/Shepherd-Go/Back-Nlj-Internal.git/entity"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/utils"
	"github.com/labstack/echo/v4"
)

type TokenMiddleware interface {
	Employee(next echo.HandlerFunc) echo.HandlerFunc
	Administrator(next echo.HandlerFunc) echo.HandlerFunc
}

type tokenMiddleware struct {
	jwt utils.JWT
}

func NewJwtMiddleware(jwt utils.JWT) TokenMiddleware {
	return &tokenMiddleware{jwt}
}

func (e *tokenMiddleware) Employee(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		ctx := c.Request().Context()

		cookie, err := c.Cookie("Authorization")
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, entity.Response{Message: "unauthorized"})
		}

		claims, err := e.jwt.PaserLoginJWT(cookie.Value)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, entity.Response{Message: err.Error()})
		}

		ctx = context.WithValue(ctx, "id", claims["id"])
		ctx = context.WithValue(ctx, "permissions", claims["permissions"])

		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)

	}
}

func (e *tokenMiddleware) Administrator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		ctx := c.Request().Context()

		permissions := ctx.Value("permissions")

		if permissions != "administrador" {
			return echo.NewHTTPError(http.StatusUnauthorized, entity.Response{Message: "unauthorized"})
		}

		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
