package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hamasfaa/project-evermos/configuration"
	"github.com/hamasfaa/project-evermos/middleware"
	"github.com/hamasfaa/project-evermos/model"
	"github.com/hamasfaa/project-evermos/service"
	"gorm.io/gorm"
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
	app.Get("/api/v1/product", middleware.AuthenticateJWT(false, controller.Config), controller.GetAllProducts)
	app.Get("/api/v1/product/:id", middleware.AuthenticateJWT(false, controller.Config), controller.GetProductByID)
	app.Delete("/api/v1/product/:id", middleware.AuthenticateJWT(false, controller.Config), controller.DeleteProduct)
	app.Put("/api/v1/product/:id", middleware.AuthenticateJWT(false, controller.Config), controller.UpdateProduct)
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

	slug := c.FormValue("slug")

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
		Slug:          slug,
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

func (controller ProductController) GetAllProducts(c *fiber.Ctx) error {
	pageStr := c.Query("page", "1")
	limitStr := c.Query("limit", "10")
	namaProduk := c.Query("nama_produk", "")
	kategoriIDStr := c.Query("category_id", "")
	tokoIDStr := c.Query("toko_id", "")
	maxHarga := c.Query("max_harga", "")
	minHarga := c.Query("min_harga", "")

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

	kategoriID, _ := strconv.Atoi(kategoriIDStr)
	tokoID, _ := strconv.Atoi(tokoIDStr)

	filterRequest := model.FilterProdukModel{
		Page:       page,
		Limit:      limit,
		Nama:       namaProduk,
		KategoriID: kategoriID,
		TokoID:     tokoID,
		MaxHarga:   maxHarga,
		MinHarga:   minHarga,
	}

	products, err := controller.ProductService.GetAllProducts(c.Context(), filterRequest)
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
		"data":    products,
	})
}

func (controller ProductController) GetProductByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Product ID is required",
			"errors":  []string{"Product ID tidak boleh kosong"},
			"data":    nil,
		})
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid Product ID",
			"errors":  []string{"Product ID harus berupa angka positif"},
			"data":    nil,
		})
	}

	product, err := controller.ProductService.GetProductByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to GET product",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	if product == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Product not found",
			"errors":  []string{"Produk tidak ditemukan"},
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to GET product",
		"errors":  nil,
		"data":    product,
	})
}

func (controller ProductController) DeleteProduct(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Product ID is required",
			"errors":  []string{"Product ID tidak boleh kosong"},
			"data":    nil,
		})
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid Product ID",
			"errors":  []string{"Product ID harus berupa angka positif"},
			"data":    nil,
		})
	}

	if err := controller.ProductService.DeleteProduct(c.Context(), id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  false,
				"message": "Product not found",
				"errors":  []string{"Produk tidak ditemukan"},
				"data":    nil,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to DELETE product",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to DELETE product",
		"errors":  nil,
		"data":    nil,
	})
}

func (controller ProductController) UpdateProduct(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Product ID is required",
			"errors":  []string{"Product ID tidak boleh kosong"},
			"data":    nil,
		})
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid Product ID",
			"errors":  []string{"Product ID harus berupa angka positif"},
			"data":    nil,
		})
	}

	namaProduk := c.FormValue("nama_produk")
	if namaProduk == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to UPDATE product",
			"errors":  []string{"Nama produk tidak boleh kosong"},
			"data":    nil,
		})
	}

	kategoriIDStr := c.FormValue("category_id")
	if kategoriIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to UPDATE product",
			"errors":  []string{"Category ID tidak boleh kosong"},
			"data":    nil,
		})
	}

	kategoriID, _ := strconv.Atoi(kategoriIDStr)
	hargaReseller := c.FormValue("harga_reseller")
	if hargaReseller == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to UPDATE product",
			"errors":  []string{"Harga reseller tidak boleh kosong"},
			"data":    nil,
		})
	}

	hargaKonsumen := c.FormValue("harga_konsumen")
	if hargaKonsumen == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to UPDATE product",
			"errors":  []string{"Harga konsumen tidak boleh kosong"},
			"data":    nil,
		})
	}

	stokStr := c.FormValue("stok")
	if stokStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to UPDATE product",
			"errors":  []string{"Stok tidak boleh kosong"},
			"data":    nil,
		})
	}
	stok, _ := strconv.Atoi(stokStr)

	deskripsi := c.FormValue("deskripsi")
	if deskripsi == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to UPDATE product",
			"errors":  []string{"Deskripsi tidak boleh kosong"},
			"data":    nil,
		})
	}

	slug := c.FormValue("slug")

	var photoUrls []string
	photos, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to UPDATE product",
			"errors":  []string{"Gagal membaca file upload"},
			"data":    nil,
		})
	}

	files := photos.File["photos"]
	if len(files) > 0 {
		for _, file := range files {
			if !controller.FileService.ValidateImageType(file.Header.Get("Content-Type")) {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"status":  false,
					"message": "Failed to UPDATE product",
					"errors":  []string{"Tipe file tidak valid"},
					"data":    nil,
				})
			}

			url, err := controller.FileService.UploadImage(file, "./uploads/produk")
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":  false,
					"message": "Failed to UPDATE product",
					"errors":  []string{err.Error()},
					"data":    nil,
				})
			}
			photoUrls = append(photoUrls, url)
		}
	}

	updateData := model.CreateProduct{
		Name:          namaProduk,
		KategoriID:    kategoriID,
		HargaReseller: hargaReseller,
		HargaKonsumen: hargaKonsumen,
		Stok:          stok,
		Deskripsi:     deskripsi,
		Url:           photoUrls,
		Slug:          slug,
	}

	if err := controller.ProductService.UpdateProduct(c.Context(), id, &updateData); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  false,
				"message": "Failed to UPDATE product",
				"errors":  []string{"Produk tidak ditemukan"},
				"data":    nil,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to UPDATE product",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to UPDATE product",
		"errors":  nil,
		"data":    nil,
	})
}
