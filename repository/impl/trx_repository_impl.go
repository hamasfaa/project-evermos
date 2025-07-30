package impl

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/repository"
	"gorm.io/gorm"
)

func NewTrxRepository(DB *gorm.DB) repository.TrxRepository {
	return &trxRepositoryImpl{DB: DB}
}

type trxRepositoryImpl struct {
	DB *gorm.DB
}

func (repo *trxRepositoryImpl) CreateTransaction(ctx context.Context, transaction *entity.Trx) (int, error) {
	err := repo.DB.WithContext(ctx).Create(transaction).Error
	if err != nil {
		return 0, err
	}
	return transaction.ID, nil
}

func (repo *trxRepositoryImpl) CreateDetailTransaction(ctx context.Context, detail []entity.DetailTrx) (int, error) {
	err := repo.DB.WithContext(ctx).Create(&detail).Error
	if err != nil {
		return 0, err
	}
	return len(detail), nil
}
