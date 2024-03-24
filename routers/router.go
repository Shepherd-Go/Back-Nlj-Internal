package routers

import (
	"net/http"

	"github.com/BBCompanyca/Back-Nlj-Internal.git/routers/groups"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	server   *echo.Echo
	health   groups.Health
	employee groups.Employee
}

func NewRouter(server *echo.Echo, health groups.Health, employee groups.Employee) *Router {
	return &Router{
		server,
		health,
		employee}
}

func (rtr *Router) Init() {
	rtr.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "remote_ip=${remote_ip}, method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
	}))

	rtr.server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	rtr.server.Use(middleware.Recover())

	basePath := rtr.server.Group("/api")

	rtr.health.Resource(basePath)
	rtr.employee.Resource(basePath)

}
