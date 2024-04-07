package groups

import (
	"github.com/BBCompanyca/Back-Nlj-Internal.git/controllers"
	"github.com/labstack/echo/v4"
)

type Employee interface {
	Resource(g *echo.Group)
}

type employee struct {
	employeeHand controllers.Employee
}

func NewGroupEmployee(employeeHand controllers.Employee) Employee {
	return &employee{employeeHand}
}

func (e *employee) Resource(g *echo.Group) {

	groupPath := g.Group("/employee")

	groupPath.POST("/create", e.employeeHand.CreateEmployee)
	groupPath.GET("/all", e.employeeHand.GetEmployees)
	groupPath.PUT("/update", e.employeeHand.UpdateEmployee)
	groupPath.DELETE("/delete", e.employeeHand.DeleteEmployee)

}
