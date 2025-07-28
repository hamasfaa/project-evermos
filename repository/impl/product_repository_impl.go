package impl

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/repository"
	"gorm.io/gorm"
)

func NewProductRepositoryImpl(DB *gorm.DB) repository.ProductRepository {
	return &productRepositoryImpl{DB: DB}
}

type productRepositoryImpl struct {
	DB *gorm.DB
}

func (repo *productRepositoryImpl) Create(ctx context.Context, productData *entity.Produk, tokoID int) error {

	if err := repo.DB.WithContext(ctx).Create(productData).Error; err != nil {
		return err
	}

	return nil
}

func (repo *productRepositoryImpl) CreateFoto(ctx context.Context, productID int, photoUrls []string) error {
	for _, url := range photoUrls {
		foto := &entity.FotoProduk{
			ProdukID: productID,
			Url:      url,
		}
		if err := repo.DB.WithContext(ctx).Create(foto).Error; err != nil {
			return err
		}
	}
	return nil
}
