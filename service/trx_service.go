package service

import (
	"context"

	"github.com/hamasfaa/project-evermos/model"
)

type TrxService interface {
	CreateTransaction(ctx context.Context, userID int, transaction *model.Transaksi) error
	GetTransactionsByUserID(ctx context.Context, userID int) (*model.AllTrx, error)
}
