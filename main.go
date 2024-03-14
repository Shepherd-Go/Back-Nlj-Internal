package main

import (
	"fmt"
	"log"

	"github.com/BBCompanyca/Back-Nlj-Internal.git/providers"
	"github.com/labstack/echo/v4"
)

func main() {

	container := providers.BuildContainer()

	if err := container.Invoke(func(server *echo.Echo) {
		server.Logger.Fatal(server.Start(fmt.Sprintf(":%d", 3000)))
	}); err != nil {
		log.Fatal(err)
	}

}
