package impl

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/repository"
	"github.com/hamasfaa/project-evermos/service"
)

func NewKategoriServiceImpl(kategoriRepository *repository.KategoriRepository) service.KategoriService {
	return &kategoriServiceImpl{KategoriRepository: *kategoriRepository}
}

type kategoriServiceImpl struct {
	repository.KategoriRepository
}

func (service *kategoriServiceImpl) CreateKategori(ctx context.Context, kategori *entity.Kategori) error {
	err := service.KategoriRepository.CreateKategori(ctx, kategori)
	if err != nil {
		return err
	}
	return nil
}
