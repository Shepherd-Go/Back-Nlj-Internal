package groups

import (
	"github.com/BBCompanyca/Back-Nlj-Internal.git/controller"
	"github.com/labstack/echo/v4"
)

type Employee interface {
	Resource(g *echo.Group)
}

type employee struct {
	employeeHand controller.Employee
}

func NewGroupEmployee(employeeHand controller.Employee) Employee {
	return &employee{employeeHand}
}

func (e *employee) Resource(g *echo.Group) {

	groupPath := g.Group("/employee")

	groupPath.POST("", e.employeeHand.CreateEmployee)

}
