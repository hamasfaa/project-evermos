package repository

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
)

type ProductRepository interface {
	Create(ctx context.Context, productData *entity.Produk, tokoID int) error
	CreateFoto(ctx context.Context, productID int, photoUrls []string) error
}
