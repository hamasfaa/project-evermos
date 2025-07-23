package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hamasfaa/project-evermos/configuration"
	"github.com/hamasfaa/project-evermos/middleware"
	"github.com/hamasfaa/project-evermos/service"
)

func NewTokoController(tokoService *service.TokoService, config configuration.Config) TokoController {
	return TokoController{TokoService: *tokoService, Config: config}
}

type TokoController struct {
	TokoService service.TokoService
	Config      configuration.Config
}

func (controller *TokoController) Route(app *fiber.App) {
	app.Get("/api/v1/toko/my", middleware.AuthenticateJWT(false, controller.Config), controller.MyToko)
}

func (controller *TokoController) MyToko(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)
	toko, err := controller.TokoService.GetMyToko(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to GET data",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}
	if toko == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to GET data",
			"errors":  []string{"Toko not found for the user"},
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to GET data",
		"error":   nil,
		"data":    toko,
	})
}
