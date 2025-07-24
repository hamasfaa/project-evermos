package impl

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/repository"
	"gorm.io/gorm"
)

func NewAlamatRepositoryImpl(DB *gorm.DB) repository.AlamatRepository {
	return &alamatRepositoryImpl{DB: DB}
}

type alamatRepositoryImpl struct {
	DB *gorm.DB
}

func (alamatRepository *alamatRepositoryImpl) Create(ctx context.Context, alamatData *entity.Alamat) error {
	err := alamatRepository.DB.WithContext(ctx).Create(alamatData).Error

	if err != nil {
		return err
	}
	return nil
}

func (alamatRepository *alamatRepositoryImpl) GetAlamatByUserID(ctx context.Context, userID int) ([]entity.Alamat, error) {
	var alamatList []entity.Alamat
	err := alamatRepository.DB.WithContext(ctx).Where("id_user = ?", userID).Find(&alamatList).Error

	if err != nil {
		return nil, err
	}

	return alamatList, nil
}
