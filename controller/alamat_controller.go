package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hamasfaa/project-evermos/configuration"
	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/middleware"
	"github.com/hamasfaa/project-evermos/model"
	"github.com/hamasfaa/project-evermos/service"
	"gorm.io/gorm"
)

func NewAlamatController(alamatService *service.AlamatService, config configuration.Config) *AlamatController {
	return &AlamatController{alamatService: *alamatService, Config: config}
}

type AlamatController struct {
	alamatService service.AlamatService
	Config        configuration.Config
}

func (controller AlamatController) Route(app *fiber.App) {
	app.Post("/api/v1/user/alamat", middleware.AuthenticateJWT(false, controller.Config), controller.CreateAlamat)
	app.Get("/api/v1/user/alamat", middleware.AuthenticateJWT(false, controller.Config), controller.GetAlamatByUserID)
	app.Get("/api/v1/user/alamat/:id", middleware.AuthenticateJWT(false, controller.Config), controller.GetAlamatByID)
	app.Delete("/api/v1/user/alamat/:id", middleware.AuthenticateJWT(false, controller.Config), controller.DeleteAlamatByID)
}

func (controller AlamatController) CreateAlamat(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)
	var alamatData model.AlamatModel
	if err := c.BodyParser(&alamatData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	alamatEntity := &entity.Alamat{
		JudulAlamat:  alamatData.JudulAlamat,
		NamaPenerima: alamatData.NamaPenerima,
		Notelp:       alamatData.NoTelp,
		DetailAlamat: alamatData.DetailAlamat,
	}

	if err := controller.alamatService.Create(c.Context(), userID, alamatEntity); err != nil {
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
		"data":    alamatData,
	})
}

func (controller AlamatController) GetAlamatByUserID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	alamatList, err := controller.alamatService.GetAlamatByUserID(c.Context(), userID)
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
		"data":    alamatList,
	})
}

func (controller AlamatController) GetAlamatByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, _ := strconv.Atoi(idStr)
	userID := c.Locals("userID").(int)

	alamat, err := controller.alamatService.GetAlamatByID(c.Context(), id, userID)
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
		"data":    alamat,
	})
}

func (controller AlamatController) DeleteAlamatByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, _ := strconv.Atoi(idStr)
	userID := c.Locals("userID").(int)

	if err := controller.alamatService.DeleteAlamatByID(c.Context(), id, userID); err != nil {

		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  false,
				"message": "Failed to DELETE data",
				"errors":  []string{err.Error()},
				"data":    nil,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to DELETE data",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to DELETE data",
		"errors":  nil,
		"data":    nil,
	})
}
