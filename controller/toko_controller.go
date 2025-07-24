package controller

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hamasfaa/project-evermos/configuration"
	"github.com/hamasfaa/project-evermos/middleware"
	"github.com/hamasfaa/project-evermos/model"
	"github.com/hamasfaa/project-evermos/service"
	"gorm.io/gorm"
)

func NewTokoController(tokoService *service.TokoService, fileService *service.FileService, config configuration.Config) *TokoController {
	return &TokoController{TokoService: *tokoService, FileService: *fileService, Config: config}
}

type TokoController struct {
	TokoService service.TokoService
	FileService service.FileService
	Config      configuration.Config
}

func (controller TokoController) Route(app *fiber.App) {
	app.Get("/api/v1/toko/my", middleware.AuthenticateJWT(false, controller.Config), controller.MyToko)
	app.Get("/api/v1/toko/:id_toko", middleware.AuthenticateJWT(false, controller.Config), controller.GetTokoByID)
	app.Get("/api/v1/toko", middleware.AuthenticateJWT(false, controller.Config), controller.GetAllTokos)
	app.Put("/api/v1/toko/:id_toko", middleware.AuthenticateJWT(false, controller.Config), controller.UpdateToko)
}

func (controller TokoController) MyToko(c *fiber.Ctx) error {
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

func (controller TokoController) GetTokoByID(c *fiber.Ctx) error {
	tokoIDstr := c.Params("id_toko")
	if tokoIDstr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Bad Request",
			"errors":  []string{"Toko ID is required"},
			"data":    nil,
		})
	}
	tokoID, err := strconv.Atoi(tokoIDstr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Bad Request",
			"errors":  []string{"Invalid Toko ID format"},
			"data":    nil,
		})
	}
	toko, err := controller.TokoService.GetTokoByID(c.Context(), tokoID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  false,
				"message": "Failed to GET data",
				"errors":  []string{"Toko tidak ditemukan"},
				"data":    nil,
			})
		}

		// For other errors
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
		"error":   nil,
		"data":    toko,
	})
}

func (controller TokoController) GetAllTokos(c *fiber.Ctx) error {
	pageStr := c.Query("page", "1")
	limitStr := c.Query("limit", "10")
	nama := c.Query("nama", "")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	if limit > 100 {
		limit = 100
	}

	paginationRequest := model.FilterModel{
		Page:  page,
		Limit: limit,
		Nama:  nama,
	}

	tokos, err := controller.TokoService.GetAllTokos(c.Context(), paginationRequest)
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
		"error":   nil,
		"data":    tokos,
	})
}

func (controller TokoController) UpdateToko(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	tokoIDstr := c.Params("id_toko")
	if tokoIDstr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Bad Request",
			"errors":  []string{"Toko ID is required"},
			"data":    nil,
		})
	}
	tokoID, err := strconv.Atoi(tokoIDstr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Bad Request",
			"errors":  []string{"Invalid Toko ID format"},
			"data":    nil,
		})
	}

	namaToko := c.FormValue("nama_toko")
	if namaToko == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Bad Request",
			"errors":  []string{"Nama Toko is required"},
			"data":    nil,
		})
	}

	var photoURL string
	file, err := c.FormFile("photo")
	if err == nil && file != nil {
		if !controller.FileService.ValidateImageType(file.Header.Get("Content-Type")) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  false,
				"message": "Bad Request",
				"errors":  []string{"Invalid image type. Allowed types: jpeg, jpg, png"},
				"data":    nil,
			})
		}

		photoURL, err = controller.FileService.UploadImage(file, "./uploads/toko")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  false,
				"message": "Failed to upload file",
				"errors":  []string{err.Error()},
				"data":    nil,
			})
		}

	}

	updateData := model.CreateToko{
		NamaToko: namaToko,
		UrlFoto:  photoURL,
		UserID:   userID,
	}

	err = controller.TokoService.UpdateToko(c.Context(), updateData, tokoID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  false,
				"message": "Failed to UPDATE data",
				"errors":  []string{"Toko not found"},
				"data":    nil,
			})
		}
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
		"error":   nil,
		"data":    "Update toko succeed",
	})
}
