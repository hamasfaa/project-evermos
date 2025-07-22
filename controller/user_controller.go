package controller

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/hamasfaa/project-evermos/common"
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
	app.Post("/api/v1/auth/login", controller.LoginUser)
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

func (controller UserController) LoginUser(c *fiber.Ctx) error {
	var user model.LoginModel
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	userData, err := controller.UserService.LoginUser(c.Context(), user.Notelp, user.KataSandi)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
			"errors":  []string{"No Telp atau kata sandi salah"},
			"data":    nil,
		})
	}
	tokenJwtResult := common.GenerateToken(userData.Notelp, userData.IsAdmin, controller.Config)

	userResponse := model.UserResponse{
		Nama:         userData.Nama,
		NoTelp:       userData.Notelp,
		TanggalLahir: userData.TanggalLahir.Format("02/01/2006"),
		Tentang:      userData.Tentang,
		Pekerjaan:    userData.Pekerjaan,
		Email:        userData.Email,
		IDProvinsi:   userData.IDProvinsi,
		IDKota:       userData.IDKota,
		Token:        tokenJwtResult,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to POST data",
		"errors":  nil,
		"data":    userResponse,
	})
}
