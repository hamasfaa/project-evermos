package repository

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/model"
)

type TrxRepository interface {
	CreateTransaction(ctx context.Context, transaction *entity.Trx) (int, error)
	CreateDetailTransaction(ctx context.Context, detail []entity.DetailTrx) (int, error)
	GetTransactionsByUserID(ctx context.Context, userID int, offset int, filterRequest model.FilterTrxModel) ([]entity.Trx, int64, error)
	GetByID(ctx context.Context, id int) (*entity.Trx, error)
}
