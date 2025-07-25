package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hamasfaa/project-evermos/configuration"
	"github.com/hamasfaa/project-evermos/controller"
	"github.com/hamasfaa/project-evermos/exception"
	repository "github.com/hamasfaa/project-evermos/repository/impl"
	service "github.com/hamasfaa/project-evermos/service/impl"
)

func main() {
	// setup configuration
	config := configuration.New()
	database := configuration.NewDatabase(config)

	// repository
	userRepository := repository.NewUserRepositoryImpl(database)
	tokoRepository := repository.NewTokoRepositoryImpl(database)
	alamatRepository := repository.NewAlamatRepositoryImpl(database)
	katgoriRepository := repository.NewKategoriRepositoryImpl(database)

	// service
	userService := service.NewUserServiceImpl(&userRepository, &tokoRepository)
	tokoService := service.NewTokoServiceImpl(&tokoRepository)
	alamatService := service.NewAlamatServiceImpl(&alamatRepository)
	kategoriService := service.NewKategoriServiceImpl(&katgoriRepository)
	locationService := service.NewLocationServiceImpl()
	fileService := service.NewFileServiceImpl()

	// controller
	userController := controller.NewUserController(&userService, &locationService, config)
	tokoController := controller.NewTokoController(&tokoService, &fileService, config)
	alamatController := controller.NewAlamatController(&alamatService, config)
	kategoriController := controller.NewKategoriController(&kategoriService, config)
	locationController := controller.NewLocationController(&locationService, config)

	//setup fiber
	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(recover.New())
	app.Use(cors.New())

	userController.Route(app)
	tokoController.Route(app)
	alamatController.Route(app)
	kategoriController.Route(app)
	locationController.Route(app)

	// start app
	err := app.Listen(config.Get("SERVER.PORT"))
	exception.PanicLogging(err)
}
