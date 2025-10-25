package main

import (
	"fmt"
	"gobase/configs"
	"gobase/configs/gobaseconfiguration"
	"gobase/internal/handler/welcomehandler"
	"gobase/internal/service/welcomeservice"
	"gobase/pkg/log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.InitLogger()

	conf := configs.Get()
	// db := db.GetDatabase(conf.Database) // you can activate default: postgresql
	app := fiber.New(gobaseconfiguration.GetConfiguration())

	apiGroup := app.Group("/api")
	// Welcome
	welcomeService := welcomeservice.NewWelcomeService()
	welcomehandler.NewWelcomeHandler(apiGroup, welcomeService)

	app.Listen(
		fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port),
	)
}
