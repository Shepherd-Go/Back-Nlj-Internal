package main

import (
	"fmt"
	"log"

	"github.com/BBCompanyca/Back-Nlj-Internal.git/config"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/providers"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/routers"
	"github.com/labstack/echo/v4"
)

func main() {

	container := providers.BuildContainer()
	port := config.Environments().Server.Port

	if err := container.Invoke(func(server *echo.Echo, routers *routers.Router) {

		routers.Init()

		server.Logger.Fatal(server.Start(fmt.Sprintf(":%d", port)))
	}); err != nil {
		log.Fatal(err)
	}

}
