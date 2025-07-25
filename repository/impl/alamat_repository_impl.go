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

func (alamatRepository *alamatRepositoryImpl) GetAlamatByID(ctx context.Context, id int, userID int) (*entity.Alamat, error) {
	var alamat entity.Alamat
	err := alamatRepository.DB.WithContext(ctx).Where("id = ? AND id_user = ?", id, userID).First(&alamat).Error

	if err != nil {
		return nil, err
	}

	return &alamat, nil
}

func (alamatRepository *alamatRepositoryImpl) DeleteAlamatByID(ctx context.Context, id int, userID int) error {
	result := alamatRepository.DB.WithContext(ctx).Where("id = ? AND id_user = ?", id, userID).Delete(&entity.Alamat{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (alamatRepository *alamatRepositoryImpl) UpdateAlamatByID(ctx context.Context, id int, userID int, alamatData *entity.Alamat) error {
	result := alamatRepository.DB.WithContext(ctx).Model(&entity.Alamat{}).Where("id = ? AND id_user = ?", id, userID).Updates(alamatData)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
