package impl

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/model"
	"github.com/hamasfaa/project-evermos/repository"
	"gorm.io/gorm"
)

func NewTokoRepositoryImpl(DB *gorm.DB) repository.TokoRepository {
	return &tokoRepositoryImpl{DB: DB}
}

type tokoRepositoryImpl struct {
	DB *gorm.DB
}

func (repo *tokoRepositoryImpl) GetByUserID(ctx context.Context, userID int) (*entity.Toko, error) {
	toko := &entity.Toko{}
	err := repo.DB.WithContext(ctx).Where("id_user = ?", userID).First(toko).Error
	if err != nil {
		return nil, err
	}
	return toko, nil
}

func (repo *tokoRepositoryImpl) Create(ctx context.Context, userID int, tokoData *model.CreateToko) error {
	toko := &entity.Toko{
		UserID:   userID,
		NamaToko: tokoData.NamaToko,
		UrlFoto:  tokoData.UrlFoto,
	}

	err := repo.DB.WithContext(ctx).Create(toko).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *tokoRepositoryImpl) GetByID(ctx context.Context, tokoID int) (*entity.Toko, error) {
	toko := &entity.Toko{}
	err := repo.DB.WithContext(ctx).Where("id = ?", tokoID).First(toko).Error
	if err != nil {
		return nil, err
	}
	return toko, nil
}

func (repo *tokoRepositoryImpl) GetAll(ctx context.Context, offset int, limit int, nama string) ([]entity.Toko, int64, error) {
	var tokos []entity.Toko
	var total int64

	query := repo.DB.WithContext(ctx).Model(&entity.Toko{})

	if nama != "" {
		query = query.Where("nama_toko LIKE ?", "%"+nama+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(limit).Offset(offset).Find(&tokos).Error; err != nil {
		return nil, 0, err
	}

	return tokos, total, nil
}
