package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/model"
	"github.com/hamasfaa/project-evermos/repository"
	"github.com/hamasfaa/project-evermos/service"
)

func NewTrxServiceImpl(trxRepository *repository.TrxRepository) service.TrxService {
	return &trxServiceImpl{TrxRepository: *trxRepository}
}

type trxServiceImpl struct {
	repository.TrxRepository
}

func (s *trxServiceImpl) CreateTransaction(ctx context.Context, userID int, transaction *model.Transaksi) error {
	kodeInvoice := fmt.Sprintf("INV-%d-%d", userID, time.Now().Unix())

	entityTransaction := &entity.Trx{
		UserID:   userID,
		AlamatID: transaction.AlamatKirim,
		Kode:     kodeInvoice,
		Metode:   transaction.Metode,
	}

	id, err := s.TrxRepository.CreateTransaction(ctx, entityTransaction)
	if err != nil {
		return err
	}

	entityTransaction.DetailTrx = make([]entity.DetailTrx, len(transaction.DetailTrx))

	for i, detail := range transaction.DetailTrx {
		entityTransaction.DetailTrx[i] = entity.DetailTrx{
			TrxID:     id,
			ProdukID:  detail.ProductID,
			Kuantitas: detail.Kuantitas,
		}
	}

	_, err = s.TrxRepository.CreateDetailTransaction(ctx, entityTransaction.DetailTrx)
	if err != nil {
		return err
	}

	return nil
}
