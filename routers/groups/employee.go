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
	jwt          middleware.TokenMiddleware
	employeeHand controllers.Employee
}

func NewGroupEmployee(jwt middleware.TokenMiddleware, employeeHand controllers.Employee) Employee {
	return &employee{jwt, employeeHand}
}

func (e *employee) Resource(g *echo.Group) {

	groupPath := g.Group("/employee")

	groupPath.POST("/create", e.employeeHand.CreateEmployee, e.jwt.Employee)
	groupPath.GET("/all", e.employeeHand.GetEmployees, e.jwt.Employee)
	groupPath.PUT("/update", e.employeeHand.UpdateEmployee)
	groupPath.DELETE("/delete", e.employeeHand.DeleteEmployee)

}
