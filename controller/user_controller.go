package controller

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/hamasfaa/project-evermos/configuration"
	"github.com/hamasfaa/project-evermos/model"
	"github.com/hamasfaa/project-evermos/service"
	"github.com/hamasfaa/project-evermos/service/impl"
)

func NewUserController(userService *service.UserService, config configuration.Config) *UserController {
	return &UserController{UserService: *userService, Config: config}
}

type UserController struct {
	service.UserService
	configuration.Config
}

func (controller UserController) Route(app *fiber.App) {
	app.Post("/api/v1/auth/register", controller.RegisterUser)
}

func (controller UserController) RegisterUser(c *fiber.Ctx) error {
	var user model.RegisterModel
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	err := controller.UserService.RegisterUser(c.Context(), user)
	if err != nil {
		if errors.Is(err, impl.ErrInvalidDateFormat) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  false,
				"message": "Failed to POST data",
				"errors":  []string{err.Error()},
				"data":    nil,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to POST data",
		"errors":  nil,
		"data":    "Register Succeed",
	})
}
