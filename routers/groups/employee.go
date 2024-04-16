package groups

import (
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/controllers"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/middleware"
	"github.com/labstack/echo/v4"
)

type Employee interface {
	Resource(g *echo.Group)
}

type employee struct {
	middlewareJWT     middleware.TokenMiddleware
	employeeHand      controllers.Employee
	activateEmailHand controllers.ActivateEmail
}

func NewGroupEmployee(middlewareJWT middleware.TokenMiddleware, employeeHand controllers.Employee, activateEmailHand controllers.ActivateEmail) Employee {
	return &employee{middlewareJWT, employeeHand, activateEmailHand}
}

func (e *employee) Resource(g *echo.Group) {

	groupPath := g.Group("/employee")

	groupPath.POST("/create", e.employeeHand.CreateEmployee, e.middlewareJWT.Employee, e.middlewareJWT.Administrator)
	groupPath.GET("/all", e.employeeHand.GetEmployees, e.middlewareJWT.Employee, e.middlewareJWT.Administrator)
	groupPath.PUT("/update", e.employeeHand.UpdateEmployee, e.middlewareJWT.Employee, e.middlewareJWT.Administrator)
	groupPath.DELETE("/delete", e.employeeHand.DeleteEmployee, e.middlewareJWT.Employee, e.middlewareJWT.Administrator)

	groupPath.PUT("/activate-email", e.activateEmailHand.ActivateEmail, e.middlewareJWT.Employee)

}
