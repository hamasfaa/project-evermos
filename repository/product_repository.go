package repository

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
)

type ProductRepository interface {
	Create(ctx context.Context, productData *entity.Produk, tokoID int) error
	CreateFoto(ctx context.Context, productID int, photoUrls []string) error
	GetAll(ctx context.Context, offset, limit int, nama string, kategoriID, tokoID int, maxHarga, minHarga string) ([]*entity.Produk, int64, error)
	GetByID(ctx context.Context, id int) (*entity.Produk, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, productData *entity.Produk) error
	UpdateStock(ctx context.Context, id int, stok int) error
}
