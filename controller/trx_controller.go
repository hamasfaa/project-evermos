package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hamasfaa/project-evermos/configuration"
	"github.com/hamasfaa/project-evermos/middleware"
	"github.com/hamasfaa/project-evermos/model"
	"github.com/hamasfaa/project-evermos/service"
)

func NewTrxController(trxService *service.TrxService, config configuration.Config) *TrxContoller {
	return &TrxContoller{TrxService: *trxService, Config: config}
}

type TrxContoller struct {
	service.TrxService
	configuration.Config
}

func (controller TrxContoller) Route(app *fiber.App) {
	app.Post("/api/v1/trx", middleware.AuthenticateJWT(false, controller.Config), controller.CreateTransaction)
	app.Get("/api/v1/trx", middleware.AuthenticateJWT(false, controller.Config), controller.GetTransactionsByUserID)
}

func (controller TrxContoller) CreateTransaction(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)
	var trxData model.Transaksi
	if err := c.BodyParser(&trxData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	err := controller.TrxService.CreateTransaction(c.Context(), userID, &trxData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to POST data",
		"errors":  nil,
		"data":    nil,
	})
}

func (controller TrxContoller) GetTransactionsByUserID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	transactions, err := controller.TrxService.GetTransactionsByUserID(c.Context(), userID)
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
		"data":    transactions,
	})
}
