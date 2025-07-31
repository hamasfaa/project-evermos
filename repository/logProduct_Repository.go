package repository

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
)

type LogProductRepository interface {
	CreateLogProduk(ctx context.Context, logProduk *entity.LogProduk) error
}
