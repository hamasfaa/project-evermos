package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hamasfaa/project-evermos/configuration"
	"github.com/hamasfaa/project-evermos/service"
)

func NewLocationController(locationService *service.LocationService, config configuration.Config) *LocationController {
	return &LocationController{LocationService: *locationService, Config: config}
}

type LocationController struct {
	service.LocationService
	configuration.Config
}

func (controller LocationController) Route(app *fiber.App) {
	app.Get("/api/v1/provcity/listprovincies", controller.GetProvince)
	app.Get("/api/v1/provcity/listcities/:prov_id", controller.GetCities)
	app.Get("/api/v1/provcity/detailprovince/:prov_id", controller.GetDetailProvince)
	app.Get("/api/v1/provcity/detailcity/:city_id", controller.GetDetailCity)
}

func (controller LocationController) GetProvince(c *fiber.Ctx) error {
	provinces, err := controller.LocationService.GetProvince(c.Context())
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
		"data":    provinces,
	})
}

func (controller LocationController) GetCities(c *fiber.Ctx) error {
	provinceID := c.Params("prov_id")
	if provinceID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Province ID is required",
			"errors":  []string{"Province ID is required"},
			"data":    nil,
		})
	}

	cities, err := controller.LocationService.GetCities(c.Context(), provinceID)
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
		"data":    cities,
	})
}

func (controller LocationController) GetDetailProvince(c *fiber.Ctx) error {
	provinceID := c.Params("prov_id")
	if provinceID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Province ID is required",
			"errors":  []string{"Province ID is required"},
			"data":    nil,
		})
	}

	province, err := controller.LocationService.GetProvinceByID(c.Context(), provinceID)
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
		"data":    province,
	})
}

func (controller LocationController) GetDetailCity(c *fiber.Ctx) error {
	cityID := c.Params("city_id")
	if cityID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "City ID is required",
			"errors":  []string{"City ID is required"},
			"data":    nil,
		})
	}
	provinceID := cityID[:2]
	city, err := controller.LocationService.GetCityByID(c.Context(), provinceID, cityID)
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
		"data":    city,
	})
}
