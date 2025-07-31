package impl

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/repository"
	"gorm.io/gorm"
)

func NewLogProductRepository(DB *gorm.DB) repository.LogProductRepository {
	return &logProductRepositoryImpl{DB: DB}
}

type logProductRepositoryImpl struct {
	DB *gorm.DB
}

func (repo *logProductRepositoryImpl) CreateLogProduk(ctx context.Context, logProduk *entity.LogProduk) error {
	return repo.DB.WithContext(ctx).Create(logProduk).Error
}
