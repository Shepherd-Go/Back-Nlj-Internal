package providers

import (
	"github.com/BBCompanyca/Back-Nlj-Internal.git/controllers"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/db"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/db/repository"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/routers"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/routers/groups"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/services"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/utils"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {

	Container = dig.New()

	_ = Container.Provide(func() *echo.Echo {
		return echo.New()
	})

	_ = Container.Provide(db.NewPostgresConnection)

	_ = Container.Provide(routers.NewRouter)

	_ = Container.Provide(utils.NewHashPassword)
	_ = Container.Provide(utils.NewLogsError)

	_ = Container.Provide(groups.NewHealthGroups)
	_ = Container.Provide(groups.NewGroupEmployee)

	_ = Container.Provide(controllers.NewHealthController)
	_ = Container.Provide(controllers.NewEmployeeController)

	_ = Container.Provide(services.NewServiceEmployee)

	_ = Container.Provide(repository.NewRepositoryEmployee)

	return Container
}
