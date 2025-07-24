package impl

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/model"
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

func (service *kategoriServiceImpl) GetAllKategori(ctx context.Context) ([]model.KategoriResponse, error) {
	kategoris, err := service.KategoriRepository.GetAllKategori(ctx)

	var kategoriModels []model.KategoriResponse
	for _, kategori := range kategoris {
		kategoriModels = append(kategoriModels, model.KategoriResponse{
			ID:           kategori.ID,
			NamaKategori: kategori.NamaKategori,
		})
	}

	if err != nil {
		return nil, err
	}
	return kategoriModels, nil
}

func (service *kategoriServiceImpl) GetKategoriByID(ctx context.Context, id int) (*model.KategoriResponse, error) {
	kategori, err := service.KategoriRepository.GetKategoriByID(ctx, id)
	if err != nil {
		return nil, err
	}

	kategoriModel := &model.KategoriResponse{
		ID:           kategori.ID,
		NamaKategori: kategori.NamaKategori,
	}

	return kategoriModel, nil
}

func (service *kategoriServiceImpl) DeleteKategori(ctx context.Context, id int) error {
	err := service.KategoriRepository.DeleteKategori(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (service *kategoriServiceImpl) UpdateKategori(ctx context.Context, kategori *entity.Kategori, id int) error {
	err := service.KategoriRepository.UpdateKategori(ctx, kategori, id)
	if err != nil {
		return err
	}
	return nil
}
