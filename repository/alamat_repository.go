package repository

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
)

type AlamatRepository interface {
	Create(ctx context.Context, alamatData *entity.Alamat) error
	GetAlamatByUserID(ctx context.Context, userID int) ([]entity.Alamat, error)
	GetAlamatByID(ctx context.Context, id int, userID int) (*entity.Alamat, error)
}
