package configuration

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hamasfaa/project-evermos/exception"
)

func NewFiberConfiguration() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
