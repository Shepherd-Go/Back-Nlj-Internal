package groups

import (
	"github.com/BBCompanyca/Back-Nlj-Internal.git/controllers"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/middleware"
	"github.com/labstack/echo/v4"
)

type Employee interface {
	Resource(g *echo.Group)
}

type employee struct {
	middlewareJWT middleware.TokenMiddleware
	employeeHand  controllers.Employee
}

func NewGroupEmployee(middlewareJWT middleware.TokenMiddleware, employeeHand controllers.Employee) Employee {
	return &employee{middlewareJWT, employeeHand}
}

func (e *employee) Resource(g *echo.Group) {

	groupPath := g.Group("/employee")

	groupPath.POST("/create", e.employeeHand.CreateEmployee, e.middlewareJWT.Employee)
	groupPath.GET("/all", e.employeeHand.GetEmployees, e.middlewareJWT.Employee)
	groupPath.PUT("/update", e.employeeHand.UpdateEmployee, e.middlewareJWT.Employee)
	groupPath.DELETE("/delete", e.employeeHand.DeleteEmployee, e.middlewareJWT.Employee)

}
