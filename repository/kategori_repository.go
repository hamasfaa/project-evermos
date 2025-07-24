package repository

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
)

type KategoriRepository interface {
	CreateKategori(ctx context.Context, kategori *entity.Kategori) error
	GetAllKategori(ctx context.Context) ([]entity.Kategori, error)
	GetKategoriByID(ctx context.Context, id int) (*entity.Kategori, error)
	DeleteKategori(ctx context.Context, id int) error
}
