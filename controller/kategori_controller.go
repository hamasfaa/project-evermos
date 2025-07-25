package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hamasfaa/project-evermos/configuration"
	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/middleware"
	"github.com/hamasfaa/project-evermos/model"
	"github.com/hamasfaa/project-evermos/service"
	"gorm.io/gorm"
)

func NewKategoriController(kategoriService *service.KategoriService, config configuration.Config) *KategoriController {
	return &KategoriController{kategoriService: *kategoriService, Config: config}
}

type KategoriController struct {
	kategoriService service.KategoriService
	Config          configuration.Config
}

func (controller KategoriController) Route(app *fiber.App) {
	app.Post("/api/v1/category", middleware.AuthenticateJWT(true, controller.Config), controller.CreateKategori)
	app.Get("/api/v1/category", middleware.AuthenticateJWT(false, controller.Config), controller.GetAllKategori)
	app.Get("/api/v1/category/:id_kategori", middleware.AuthenticateJWT(false, controller.Config), controller.GetKategoriByID)
	app.Delete("/api/v1/category/:id_kategori", middleware.AuthenticateJWT(true, controller.Config), controller.DeleteKategori)
	app.Put("/api/v1/category/:id_kategori", middleware.AuthenticateJWT(true, controller.Config), controller.UpdateKategori)
}

func (controller KategoriController) CreateKategori(c *fiber.Ctx) error {
	var kategoriData model.Kategori
	if err := c.BodyParser(&kategoriData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	kategoriEntity := &entity.Kategori{
		NamaKategori: kategoriData.NamaKategori,
	}

	if err := controller.kategoriService.CreateKategori(c.Context(), kategoriEntity); err != nil {
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
		"data":    kategoriData,
	})
}

func (controller KategoriController) GetAllKategori(c *fiber.Ctx) error {
	kategoriModels, err := controller.kategoriService.GetAllKategori(c.Context())
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
		"data":    kategoriModels,
	})
}

func (controller KategoriController) GetKategoriByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id_kategori")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid category ID",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	kategoriModel, err := controller.kategoriService.GetKategoriByID(c.Context(), id)
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
		"data":    kategoriModel,
	})
}

func (controller KategoriController) DeleteKategori(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id_kategori")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid category ID",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	if err := controller.kategoriService.DeleteKategori(c.Context(), id); err != nil {
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

func (controller KategoriController) UpdateKategori(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id_kategori")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid category ID",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	var kategoriData model.Kategori
	if err := c.BodyParser(&kategoriData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to parse request body",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	kategoriEntity := &entity.Kategori{
		ID:           id,
		NamaKategori: kategoriData.NamaKategori,
	}

	if err := controller.kategoriService.UpdateKategori(c.Context(), kategoriEntity, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to UPDATE data",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to UPDATE data",
		"errors":  nil,
		"data":    kategoriData,
	})
}
