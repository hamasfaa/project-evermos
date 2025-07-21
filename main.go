package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hamasfaa/project-evermos/configuration"
	"github.com/hamasfaa/project-evermos/exception"
)

func main() {
	// setup configuration
	config := configuration.New()
	_ = configuration.NewDatabase(config)

	//setup fiber
	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(recover.New())
	app.Use(cors.New())

	// start app
	err := app.Listen(config.Get("SERVER.PORT"))
	exception.PanicLogging(err)
}
