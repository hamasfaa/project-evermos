package service

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
)

type KategoriService interface {
	CreateKategori(ctx context.Context, kategori *entity.Kategori) error
}
