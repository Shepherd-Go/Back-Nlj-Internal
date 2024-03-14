package routers

import (
	"github.com/BBCompanyca/Back-Nlj-Internal.git/routers/groups"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	server *echo.Echo
	health groups.Health
}

func NewRouter(server *echo.Echo, health groups.Health) *Router {
	return &Router{
		server,
		health}
}

func (rtr *Router) Init() {
	rtr.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "remote_ip=${remote_ip}, method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
	}))

	rtr.server.Use(middleware.Recover())

	basePath := rtr.server.Group("/api")

	rtr.health.Resource(basePath)

}
