package controller

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/hamasfaa/project-evermos/common"
	"github.com/hamasfaa/project-evermos/configuration"
	"github.com/hamasfaa/project-evermos/middleware"
	"github.com/hamasfaa/project-evermos/model"
	"github.com/hamasfaa/project-evermos/service"
	"github.com/hamasfaa/project-evermos/service/impl"
)

func NewUserController(userService *service.UserService, locationService *service.LocationService, config configuration.Config) *UserController {
	return &UserController{UserService: *userService, LocationService: *locationService, Config: config}
}

type UserController struct {
	service.UserService
	service.LocationService
	configuration.Config
}

func (controller UserController) Route(app *fiber.App) {
	app.Post("/api/v1/auth/register", controller.RegisterUser)
	app.Post("/api/v1/auth/login", controller.LoginUser)
	app.Get("/api/v1/user", middleware.AuthenticateJWT(false, controller.Config), controller.Me)
	app.Put("/api/v1/user", middleware.AuthenticateJWT(false, controller.Config), controller.UpdateUser)
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
	tokenJwtResult := common.GenerateToken(userData.Notelp, userData.IsAdmin, userData.ID, controller.Config)

	provinceData, err := controller.LocationService.GetProvinceByID(c.Context(), userData.IDProvinsi)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to GET data",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}
	cityData, err := controller.LocationService.GetCityByID(c.Context(), userData.IDProvinsi, userData.IDKota)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to GET data",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	userResponse := model.UserResponse{
		Nama:         userData.Nama,
		NoTelp:       userData.Notelp,
		TanggalLahir: userData.TanggalLahir.Format("02/01/2006"),
		Tentang:      userData.Tentang,
		Pekerjaan:    userData.Pekerjaan,
		Email:        userData.Email,
		IDProvinsi:   provinceData,
		IDKota:       cityData,
		Token:        tokenJwtResult,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to POST data",
		"errors":  nil,
		"data":    userResponse,
	})
}

func (controller UserController) Me(c *fiber.Ctx) error {
	noTelp, ok := c.Locals("noTelp").(string)

	if !ok || noTelp == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Unauthorized",
			"errors":  []string{"Invalid user data"},
			"data":    nil,
		})
	}

	meResponse, err := controller.UserService.Me(c.Context(), noTelp)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to GET data",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to GET data",
		"errors":  nil,
		"data":    meResponse,
	})
}

func (controller UserController) UpdateUser(c *fiber.Ctx) error {
	noTelp, ok := c.Locals("noTelp").(string)

	if !ok || noTelp == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Unauthorized",
			"errors":  []string{"Invalid user data"},
			"data":    nil,
		})
	}

	var user model.RegisterModel
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to parse request body",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	user.NoTelp = noTelp

	err := controller.UserService.UpdateUser(c.Context(), noTelp, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to update user",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "User updated successfully",
		"errors":  nil,
		"data":    nil,
	})
}
