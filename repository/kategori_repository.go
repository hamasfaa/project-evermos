package repository

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
)

type KategoriRepository interface {
	CreateKategori(ctx context.Context, kategori *entity.Kategori) error
}
