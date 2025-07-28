package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hamasfaa/project-evermos/configuration"
	"github.com/hamasfaa/project-evermos/middleware"
	"github.com/hamasfaa/project-evermos/model"
	"github.com/hamasfaa/project-evermos/service"
)

func NewProductController(productService *service.ProductService, fileService *service.FileService, config configuration.Config) *ProductController {
	return &ProductController{ProductService: *productService, FileService: *fileService, Config: config}
}

type ProductController struct {
	ProductService service.ProductService
	FileService    service.FileService
	Config         configuration.Config
}

func (controller ProductController) Route(app *fiber.App) {
	app.Post("/api/v1/product", middleware.AuthenticateJWT(false, controller.Config), controller.CreateProduct)
}

func (controller ProductController) CreateProduct(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	namaProduk := c.FormValue("nama_produk")
	if namaProduk == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
			"errors":  []string{"Nama produk tidak boleh kosong"},
			"data":    nil,
		})
	}

	kategoriIDStr := c.FormValue("category_id")
	if kategoriIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
			"errors":  []string{"Category ID tidak boleh kosong"},
			"data":    nil,
		})
	}
	kategoriID, _ := strconv.Atoi(kategoriIDStr)

	hargaReseller := c.FormValue("harga_reseller")
	if hargaReseller == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
			"errors":  []string{"Harga reseller tidak boleh kosong"},
			"data":    nil,
		})
	}

	hargaKonsumen := c.FormValue("harga_konsumen")
	if hargaKonsumen == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
			"errors":  []string{"Harga konsumen tidak boleh kosong"},
			"data":    nil,
		})
	}

	stokStr := c.FormValue("stok")
	if stokStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
			"errors":  []string{"Stok tidak boleh kosong"},
			"data":    nil,
		})
	}
	stok, _ := strconv.Atoi(stokStr)

	deskripsi := c.FormValue("deskripsi")
	if deskripsi == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
			"errors":  []string{"Deskripsi tidak boleh kosong"},
			"data":    nil,
		})
	}

	var photoUrls []string
	photos, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
			"errors":  []string{"Gagal membaca file upload"},
			"data":    nil,
		})
	}

	files := photos.File["photos"]
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
			"errors":  []string{"Minimal satu foto produk harus diunggah"},
			"data":    nil,
		})
	}

	for _, file := range files {
		if !controller.FileService.ValidateImageType(file.Header.Get("Content-Type")) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  false,
				"message": "Failed to POST data",
				"errors":  []string{"Tipe file tidak valid. Hanya jpeg, jpg, png yang diperbolehkan"},
				"data":    nil,
			})
		}

		url, err := controller.FileService.UploadImage(file, "./uploads/produk")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  false,
				"message": "Failed to POST data",
				"errors":  []string{err.Error()},
				"data":    nil,
			})
		}
		photoUrls = append(photoUrls, url)
	}

	createData := model.CreateProduct{
		Name:          namaProduk,
		KategoriID:    kategoriID,
		HargaReseller: hargaReseller,
		HargaKonsumen: hargaKonsumen,
		Stok:          stok,
		Deskripsi:     deskripsi,
		Url:           photoUrls,
	}

	if err := controller.ProductService.CreateProduct(c.Context(), userID, &createData); err != nil {
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
		"data":    nil,
	})
}
