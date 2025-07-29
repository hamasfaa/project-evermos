package impl

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/model"
	"github.com/hamasfaa/project-evermos/repository"
	"github.com/hamasfaa/project-evermos/service"
)

func NewProductServiceImpl(productRepository *repository.ProductRepository, tokoRepository *repository.TokoRepository) service.ProductService {
	return &productServiceImpl{ProductRepository: *productRepository, TokoRepository: *tokoRepository}
}

type productServiceImpl struct {
	repository.ProductRepository
	repository.TokoRepository
}

func (service *productServiceImpl) CreateProduct(ctx context.Context, userID int, productData *model.CreateProduct) error {
	toko, err := service.TokoRepository.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}
	tokoID := toko.ID

	product := &entity.Produk{
		NamaProduk:    productData.Name,
		HargaReseller: productData.HargaReseller,
		HargaKonsumen: productData.HargaKonsumen,
		Stok:          productData.Stok,
		Deskripsi:     productData.Deskripsi,
		KategoriID:    productData.KategoriID,
		TokoID:        tokoID,
		Slug:          productData.Slug,
	}

	if err = service.ProductRepository.Create(ctx, product, tokoID); err != nil {
		return err
	}

	if len(productData.Url) > 0 {
		if err = service.ProductRepository.CreateFoto(ctx, product.ID, productData.Url); err != nil {
			return err
		}
	}

	return nil
}

func (service *productServiceImpl) GetAllProducts(ctx context.Context, pagination model.FilterProdukModel) (*model.AllProduk, error) {
	offset := (pagination.Page - 1) * pagination.Limit

	products, total, err := service.ProductRepository.GetAll(ctx, offset, pagination.Limit, pagination.Nama, pagination.KategoriID, pagination.TokoID, pagination.MaxHarga, pagination.MinHarga)
	if err != nil {
		return nil, err
	}
	var productModels []model.Produk
	for _, p := range products {
		var fotoResponses []model.FotoProdukResponse
		for _, foto := range p.FotoProduks {
			fotoResponses = append(fotoResponses, model.FotoProdukResponse{
				ID:       foto.ID,
				ProdukID: foto.ProdukID,
				Url:      foto.Url,
			})
		}

		productModels = append(productModels, model.Produk{
			ID:            p.ID,
			NamaProduk:    p.NamaProduk,
			Slug:          p.Slug,
			HargaReseller: p.HargaReseller,
			HargaKonsumen: p.HargaKonsumen,
			Stok:          p.Stok,
			Deskripsi:     p.Deskripsi,
			Toko:          model.TokoModel{ID: p.Toko.ID, NamaToko: p.Toko.NamaToko, UrlFoto: p.Toko.UrlFoto},
			Kategori:      model.KategoriResponse{ID: p.Kategori.ID, NamaKategori: p.Kategori.NamaKategori},
			Foto:          fotoResponses,
		})
	}

	totalPages := (len(products) + pagination.Limit - 1) / pagination.Limit
	hasNext := pagination.Page < totalPages
	hasPrev := pagination.Page > 1

	result := &model.AllProduk{
		Data:       productModels,
		TotalItems: total,
		TotalPages: totalPages,
		Page:       pagination.Page,
		Limit:      pagination.Limit,
		HasNext:    hasNext,
		HasPrev:    hasPrev,
	}
	return result, nil
}

func (service *productServiceImpl) GetProductByID(ctx context.Context, id int) (*model.Produk, error) {
	product, err := service.ProductRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	var fotoResponses []model.FotoProdukResponse
	for _, foto := range product.FotoProduks {
		fotoResponses = append(fotoResponses, model.FotoProdukResponse{
			ID:       foto.ID,
			ProdukID: foto.ProdukID,
			Url:      foto.Url,
		})
	}

	result := &model.Produk{
		ID:            product.ID,
		NamaProduk:    product.NamaProduk,
		Slug:          product.Slug,
		HargaReseller: product.HargaReseller,
		HargaKonsumen: product.HargaKonsumen,
		Stok:          product.Stok,
		Deskripsi:     product.Deskripsi,
		Toko:          model.TokoModel{ID: product.Toko.ID, NamaToko: product.Toko.NamaToko, UrlFoto: product.Toko.UrlFoto},
		Kategori:      model.KategoriResponse{ID: product.Kategori.ID, NamaKategori: product.Kategori.NamaKategori},
		Foto:          fotoResponses,
	}

	return result, nil
}

func (service *productServiceImpl) DeleteProduct(ctx context.Context, id int) error {
	if err := service.ProductRepository.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (service *productServiceImpl) UpdateProduct(ctx context.Context, id int, productData *model.CreateProduct) error {
	product := &entity.Produk{
		ID:            id,
		NamaProduk:    productData.Name,
		HargaReseller: productData.HargaReseller,
		HargaKonsumen: productData.HargaKonsumen,
		Stok:          productData.Stok,
		Deskripsi:     productData.Deskripsi,
		KategoriID:    productData.KategoriID,
		Slug:          productData.Slug,
	}

	if err := service.ProductRepository.Update(ctx, id, product); err != nil {
		return err
	}

	if len(productData.Url) > 0 {
		if err := service.ProductRepository.CreateFoto(ctx, id, productData.Url); err != nil {
			return err
		}
	}

	return nil
}
