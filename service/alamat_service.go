package service

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/model"
)

type AlamatService interface {
	Create(ctx context.Context, userID int, alamatData *entity.Alamat) error
	GetAlamatByUserID(ctx context.Context, userID int) ([]model.AlamatResponse, error)
	GetAlamatByID(ctx context.Context, id int, userID int) (*model.AlamatResponse, error)
	DeleteAlamatByID(ctx context.Context, id int, userID int) error
	UpdateAlamatByID(ctx context.Context, id int, userID int, alamatData *model.UpdateAlamatModel) error
}
