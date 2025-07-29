package service

import (
	"context"

	"github.com/hamasfaa/project-evermos/model"
)

type ProductService interface {
	CreateProduct(ctx context.Context, userID int, productData *model.CreateProduct) error
	GetAllProducts(ctx context.Context, pagination model.FilterProdukModel) (*model.AllProduk, error)
}
