package service

import (
	"context"

	"github.com/hamasfaa/project-evermos/model"
)

type TrxService interface {
	CreateTransaction(ctx context.Context, userID int, transaction *model.Transaksi) error
	GetTransactionsByUserID(ctx context.Context, userID int, filterRequest model.FilterTrxModel) (*model.AllTrx, error)
	GetTransactionByID(ctx context.Context, id int) (*model.Trx, error)
}
