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
