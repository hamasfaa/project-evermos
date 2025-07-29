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

func (repo *productRepositoryImpl) GetAll(ctx context.Context, offset, limit int, nama string, kategoriID, tokoID int, maxHarga, minHarga string) ([]*entity.Produk, int64, error) {
	var products []*entity.Produk
	var total int64

	query := repo.DB.WithContext(ctx).Model(&entity.Produk{})
	if nama != "" {
		query = query.Where("nama_produk LIKE ?", "%"+nama+"%")
	}
	if kategoriID > 0 {
		query = query.Where("id_kategori = ?", kategoriID)
	}
	if tokoID > 0 {
		query = query.Where("id_toko = ?", tokoID)
	}
	if maxHarga != "" {
		query = query.Where("harga_konsumen <= ?", maxHarga)
	}
	if minHarga != "" {
		query = query.Where("harga_konsumen >= ?", minHarga)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	query = query.Preload("Toko").Preload("Kategori").Preload("FotoProduks")
	if err := query.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (repo *productRepositoryImpl) GetByID(ctx context.Context, id int) (*entity.Produk, error) {
	var product entity.Produk
	err := repo.DB.WithContext(ctx).Preload("Toko").Preload("Kategori").Preload("FotoProduks").First(&product, id).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *productRepositoryImpl) Delete(ctx context.Context, id int) error {
	result := repo.DB.WithContext(ctx).Where("id = ?", id).Delete(&entity.Produk{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
