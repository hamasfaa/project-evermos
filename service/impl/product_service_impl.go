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
