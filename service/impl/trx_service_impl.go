package impl

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/model"
	"github.com/hamasfaa/project-evermos/repository"
	"github.com/hamasfaa/project-evermos/service"
)

func NewTrxServiceImpl(trxRepository *repository.TrxRepository, productRepository *repository.ProductRepository) service.TrxService {
	return &trxServiceImpl{TrxRepository: *trxRepository, ProductRepository: *productRepository}
}

type trxServiceImpl struct {
	repository.TrxRepository
	repository.ProductRepository
}

func (s *trxServiceImpl) CreateTransaction(ctx context.Context, userID int, transaction *model.Transaksi) error {
	productData := make(map[int]*entity.Produk)
	var totalHarga int

	for _, detail := range transaction.DetailTrx {
		if _, exist := productData[detail.ProductID]; !exist {
			product, err := s.ProductRepository.GetByID(ctx, detail.ProductID)
			if err != nil {
				return err
			}
			productData[detail.ProductID] = product
		}

		product := productData[detail.ProductID]

		if product.Stok < detail.Kuantitas {
			return fmt.Errorf("insufficient stock for product %s. Available: %d, Requested: %d",
				product.NamaProduk, product.Stok, detail.Kuantitas)
		}

		hargaKonsumen, _ := strconv.Atoi(product.HargaKonsumen)

		totalHarga += hargaKonsumen * detail.Kuantitas
	}

	kodeInvoice := fmt.Sprintf("INV-%d-%d", userID, time.Now().Unix())

	entityTransaction := &entity.Trx{
		UserID:     userID,
		AlamatID:   transaction.AlamatKirim,
		Kode:       kodeInvoice,
		Metode:     transaction.Metode,
		HargaTotal: totalHarga,
	}

	id, err := s.TrxRepository.CreateTransaction(ctx, entityTransaction)
	if err != nil {
		return err
	}

	entityTransaction.DetailTrx = make([]entity.DetailTrx, len(transaction.DetailTrx))

	for i, detail := range transaction.DetailTrx {
		product := productData[detail.ProductID]
		hargaKonsumen, _ := strconv.Atoi(product.HargaKonsumen)

		hargaTotal := hargaKonsumen * detail.Kuantitas

		entityTransaction.DetailTrx[i] = entity.DetailTrx{
			TrxID:      id,
			ProdukID:   detail.ProductID,
			Kuantitas:  detail.Kuantitas,
			HargaTotal: hargaTotal,
		}
	}

	_, err = s.TrxRepository.CreateDetailTransaction(ctx, entityTransaction.DetailTrx)
	if err != nil {
		return err
	}

	for _, detail := range entityTransaction.DetailTrx {
		product := productData[detail.ProdukID]
		product.Stok -= detail.Kuantitas
		err := s.ProductRepository.UpdateStock(ctx, product.ID, product.Stok)
		if err != nil {
			return fmt.Errorf("failed to update stock for product %s: %w", product.NamaProduk, err)
		}
	}

	return nil
}
