package repository

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
)

type TrxRepository interface {
	CreateTransaction(ctx context.Context, transaction *entity.Trx) (int, error)
	CreateDetailTransaction(ctx context.Context, detail []entity.DetailTrx) (int, error)
	GetTransactionsByUserID(ctx context.Context, userID int) ([]entity.Trx, error)
}
