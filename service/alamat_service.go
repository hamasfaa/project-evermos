package service

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/model"
)

type AlamatService interface {
	Create(ctx context.Context, userID int, alamatData *entity.Alamat) error
	GetAlamatByUserID(ctx context.Context, userID int) ([]model.AlamatResponse, error)
}
