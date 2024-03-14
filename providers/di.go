package providers

import (
	"github.com/BBCompanyca/Back-Nlj-Internal.git/controller"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/routers"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/routers/groups"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {

	Container = dig.New()

	_ = Container.Provide(func() *echo.Echo {
		return echo.New()
	})

	_ = Container.Provide(routers.NewRouter)

	_ = Container.Provide(groups.NewHealthGroups)

	_ = Container.Provide(controller.NewHealthController)

	return Container
}
