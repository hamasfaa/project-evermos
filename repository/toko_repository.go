package repository

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/model"
)

type TokoRepository interface {
	GetByUserID(ctx context.Context, userID int) (*entity.Toko, error)
	Create(ctx context.Context, userID int, tokoData *model.CreateToko) error
	GetByID(ctx context.Context, tokoID int) (*entity.Toko, error)
	GetAll(ctx context.Context, offset int, limit int, nama string) ([]entity.Toko, int64, error)
	Update(ctx context.Context, tokoData *model.CreateToko, tokoID int) error
}
