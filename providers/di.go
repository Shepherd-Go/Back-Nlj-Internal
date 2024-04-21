package providers

import (
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/controllers"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/db"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/db/repository"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/middleware"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/routers"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/routers/groups"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/services"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/utils"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {

	Container = dig.New()

	_ = Container.Provide(func() *echo.Echo {
		return echo.New()
	})

	_ = Container.Provide(middleware.NewJwtMiddleware)

	_ = Container.Provide(routers.NewRouter)

	_ = Container.Provide(db.NewPostgresConnection)

	_ = Container.Provide(utils.NewHashPassword)
	_ = Container.Provide(utils.NewJWT)
	_ = Container.Provide(utils.NewSendEmail)

	_ = Container.Provide(groups.NewHealthGroups)
	_ = Container.Provide(groups.NewGroupSession)
	_ = Container.Provide(groups.NewGroupEmployee)

	_ = Container.Provide(controllers.NewHealthController)
	_ = Container.Provide(controllers.NewSessionController)
	_ = Container.Provide(controllers.NewEmployeeController)
	_ = Container.Provide(controllers.NewActiveEmailControler)

	_ = Container.Provide(services.NewSessionService)
	_ = Container.Provide(services.NewServiceEmployee)
	_ = Container.Provide(services.NewActivateEmailService)

	_ = Container.Provide(repository.NewRepositoryEmployee)
	_ = Container.Provide(repository.NewLogInRepository)

	return Container
}
