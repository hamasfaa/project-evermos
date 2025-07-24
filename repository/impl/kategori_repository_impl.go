package impl

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/repository"
	"gorm.io/gorm"
)

func NewKategoriRepositoryImpl(DB *gorm.DB) repository.KategoriRepository {
	return &kategoriRepositoryImpl{DB: DB}
}

type kategoriRepositoryImpl struct {
	DB *gorm.DB
}

func (kategoriRepository *kategoriRepositoryImpl) CreateKategori(ctx context.Context, kategori *entity.Kategori) error {
	err := kategoriRepository.DB.WithContext(ctx).Create(kategori).Error
	if err != nil {
		return err
	}
	return nil
}

func (kategoriRepository *kategoriRepositoryImpl) GetAllKategori(ctx context.Context) ([]entity.Kategori, error) {
	var kategoris []entity.Kategori
	err := kategoriRepository.DB.WithContext(ctx).Find(&kategoris).Error
	if err != nil {
		return nil, err
	}
	return kategoris, nil
}

func (kategoriRepository *kategoriRepositoryImpl) GetKategoriByID(ctx context.Context, id int) (*entity.Kategori, error) {
	var kategori entity.Kategori
	err := kategoriRepository.DB.WithContext(ctx).Where("id = ?", id).First(&kategori).Error
	if err != nil {
		return nil, err
	}
	return &kategori, nil
}

func (kategoriRepository *kategoriRepositoryImpl) DeleteKategori(ctx context.Context, id int) error {
	result := kategoriRepository.DB.WithContext(ctx).Where("id = ?", id).Delete(&entity.Kategori{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (kategoriRepository *kategoriRepositoryImpl) UpdateKategori(ctx context.Context, kategori *entity.Kategori, id int) error {
	result := kategoriRepository.DB.WithContext(ctx).Model(&entity.Kategori{}).Where("id = ?", id).Updates(kategori)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
