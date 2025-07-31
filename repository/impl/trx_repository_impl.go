package impl

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/model"
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

func (repo *trxRepositoryImpl) CreateDetailTransaction(ctx context.Context, detail []entity.DetailTrx) ([]int, error) {
	err := repo.DB.WithContext(ctx).Create(&detail).Error
	if err != nil {
		return nil, err
	}
	var IDs []int
	for _, d := range detail {
		IDs = append(IDs, d.ID)
	}
	return IDs, nil
}

func (repo *trxRepositoryImpl) GetTransactionsByUserID(ctx context.Context, userID int, offset int, filterRequest model.FilterTrxModel) ([]entity.Trx, int64, error) {
	var transactions []entity.Trx
	var total int64

	query := repo.DB.WithContext(ctx).Where("id_user = ?", userID)

	if filterRequest.Search != "" {
		query = query.Where("kode_invoice LIKE ?", "%"+filterRequest.Search+"%")
	}

	if err := query.Model(&entity.Trx{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	query = query.Preload("User").Preload("Alamat").Preload("DetailTrx").Preload("DetailTrx.Produk").Preload("DetailTrx.Produk.Toko").Preload("DetailTrx.Produk.Kategori").Preload("DetailTrx.Produk.FotoProduks")

	if filterRequest.Limit > 0 {
		query = query.Limit(filterRequest.Limit)
	}

	if offset > 0 {
		query = query.Offset(offset)
	}

	if err := query.Find(&transactions).Error; err != nil {
		return nil, 0, err
	}

	return transactions, total, nil
}

func (repo *trxRepositoryImpl) GetByID(ctx context.Context, id int) (*entity.Trx, error) {
	var transaction entity.Trx
	err := repo.DB.WithContext(ctx).Where("id = ?", id).Preload("User").Preload("Alamat").Preload("DetailTrx").Preload("DetailTrx.Produk").Preload("DetailTrx.Produk.Toko").Preload("DetailTrx.Produk.Kategori").Preload("DetailTrx.Produk.FotoProduks").First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}
