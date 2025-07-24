package service

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/model"
)

type KategoriService interface {
	CreateKategori(ctx context.Context, kategori *entity.Kategori) error
	GetAllKategori(ctx context.Context) ([]model.KategoriResponse, error)
	GetKategoriByID(ctx context.Context, id int) (*model.KategoriResponse, error)
	DeleteKategori(ctx context.Context, id int) error
}
