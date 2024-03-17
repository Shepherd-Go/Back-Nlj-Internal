package groups

import (
	"github.com/BBCompanyca/Back-Nlj-Internal.git/controllers"
	"github.com/labstack/echo/v4"
)

type Health interface {
	Resource(g *echo.Group)
}

type health struct {
	healthHand controllers.Health
}

func NewHealthGroups(healthHand controllers.Health) Health {
	return &health{healthHand}
}

func (h *health) Resource(g *echo.Group) {

	groupPath := g.Group("/health")

	groupPath.GET("", h.healthHand.Health)

}
