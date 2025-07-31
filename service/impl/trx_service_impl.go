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

func (s *trxServiceImpl) GetTransactionsByUserID(ctx context.Context, userID int, filterRequest model.FilterTrxModel) (*model.AllTrx, error) {
	offset := (filterRequest.Page - 1) * filterRequest.Limit

	transactions, total, err := s.TrxRepository.GetTransactionsByUserID(ctx, userID, offset, filterRequest)
	if err != nil {
		return nil, err
	}
	var trxDetails []model.Trx
	for _, trx := range transactions {
		details := make([]model.TrxDetail, len(trx.DetailTrx))
		for i, detail := range trx.DetailTrx {
			var fotoResponses []model.FotoProdukResponse
			for _, foto := range detail.Produk.FotoProduks {
				fotoResponses = append(fotoResponses, model.FotoProdukResponse{
					ID:       foto.ID,
					ProdukID: foto.ProdukID,
					Url:      foto.Url,
				})
			}
			details[i] = model.TrxDetail{
				Produk: model.Produk{
					ID:            detail.Produk.ID,
					NamaProduk:    detail.Produk.NamaProduk,
					Slug:          detail.Produk.Slug,
					HargaReseller: detail.Produk.HargaReseller,
					HargaKonsumen: detail.Produk.HargaKonsumen,
					Deskripsi:     detail.Produk.Deskripsi,
					Toko: model.TokoModel{
						NamaToko: detail.Produk.Toko.NamaToko,
						UrlFoto:  detail.Produk.Toko.UrlFoto,
					},
					Kategori: model.KategoriResponse{
						ID:           detail.Produk.Kategori.ID,
						NamaKategori: detail.Produk.Kategori.NamaKategori,
					},
					Foto: fotoResponses,
				},
				Kuantitas:  detail.Kuantitas,
				HargaTotal: strconv.Itoa(detail.HargaTotal),
			}
		}
		trxDetails = append(trxDetails, model.Trx{
			ID:         trx.ID,
			HargaTotal: trx.HargaTotal,
			Kode:       trx.Kode,
			Metode:     trx.Metode,
			Alamat: model.AlamatResponse{
				ID:           trx.Alamat.ID,
				JudulAlamat:  trx.Alamat.JudulAlamat,
				NamaPenerima: trx.Alamat.NamaPenerima,
				NoTelp:       trx.Alamat.Notelp,
				DetailAlamat: trx.Alamat.DetailAlamat,
			},
			Detail: details,
		})
	}

	totalPages := (int(total) + filterRequest.Limit - 1) / filterRequest.Limit
	hasNext := filterRequest.Page < totalPages
	hasPrev := filterRequest.Page > 1

	result := &model.AllTrx{
		Data:       trxDetails,
		TotalItems: total,
		TotalPages: totalPages,
		Page:       filterRequest.Page,
		Limit:      filterRequest.Limit,
		HasNext:    hasNext,
		HasPrev:    hasPrev,
	}

	return result, nil
}

func (s *trxServiceImpl) GetTransactionByID(ctx context.Context, id int) (*model.Trx, error) {
	transaction, err := s.TrxRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	var details []model.TrxDetail
	for _, detail := range transaction.DetailTrx {
		var fotoResponses []model.FotoProdukResponse
		for _, foto := range detail.Produk.FotoProduks {
			fotoResponses = append(fotoResponses, model.FotoProdukResponse{
				ID:       foto.ID,
				ProdukID: foto.ProdukID,
				Url:      foto.Url,
			})
		}
		details = append(details, model.TrxDetail{
			Produk: model.Produk{
				ID:            detail.Produk.ID,
				NamaProduk:    detail.Produk.NamaProduk,
				Slug:          detail.Produk.Slug,
				HargaReseller: detail.Produk.HargaReseller,
				HargaKonsumen: detail.Produk.HargaKonsumen,
				Deskripsi:     detail.Produk.Deskripsi,
				Toko: model.TokoModel{
					NamaToko: detail.Produk.Toko.NamaToko,
					UrlFoto:  detail.Produk.Toko.UrlFoto,
				},
				Kategori: model.KategoriResponse{
					ID:           detail.Produk.Kategori.ID,
					NamaKategori: detail.Produk.Kategori.NamaKategori,
				},
				Foto: fotoResponses,
			},
			Kuantitas:  detail.Kuantitas,
			HargaTotal: strconv.Itoa(detail.HargaTotal),
		})
	}

	result := &model.Trx{
		ID:         transaction.ID,
		HargaTotal: transaction.HargaTotal,
		Kode:       transaction.Kode,
		Metode:     transaction.Metode,
		Alamat: model.AlamatResponse{
			ID:           transaction.Alamat.ID,
			JudulAlamat:  transaction.Alamat.JudulAlamat,
			NamaPenerima: transaction.Alamat.NamaPenerima,
			NoTelp:       transaction.Alamat.Notelp,
			DetailAlamat: transaction.Alamat.DetailAlamat,
		},
		Detail: details,
	}

	return result, nil
}
