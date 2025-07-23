package impl

import (
	"context"
	"math"

	"github.com/hamasfaa/project-evermos/model"
	"github.com/hamasfaa/project-evermos/repository"
	"github.com/hamasfaa/project-evermos/service"
)

func NewTokoServiceImpl(tokoRepository *repository.TokoRepository) service.TokoService {
	return &tokoServiceImpl{TokoRepository: *tokoRepository}
}

type tokoServiceImpl struct {
	repository.TokoRepository
}

func (t *tokoServiceImpl) GetMyToko(ctx context.Context, userID int) (*model.MyToko, error) {
	toko, err := t.TokoRepository.GetByUserID(ctx, userID)

	tokoData := &model.MyToko{
		ID:       toko.ID,
		NamaToko: toko.NamaToko,
		UrlFoto:  toko.UrlFoto,
		UserID:   toko.UserID,
	}

	if err != nil {
		return nil, err
	}
	return tokoData, nil
}

func (t *tokoServiceImpl) GetTokoByID(ctx context.Context, tokoID int) (*model.TokoModel, error) {
	toko, err := t.TokoRepository.GetByID(ctx, tokoID)
	if err != nil {
		return nil, err
	}

	tokoModel := &model.TokoModel{
		ID:       toko.ID,
		NamaToko: toko.NamaToko,
		UrlFoto:  toko.UrlFoto,
	}

	return tokoModel, nil
}

func (t *tokoServiceImpl) GetAllTokos(ctx context.Context, pagination model.FilterModel) (*model.AllToko, error) {
	offset := (pagination.Page - 1) * pagination.Limit

	tokos, total, err := t.TokoRepository.GetAll(ctx, offset, pagination.Limit, pagination.Nama)
	if err != nil {
		return nil, err
	}
	var tokoModels []model.TokoModel
	for _, toko := range tokos {
		tokoModels = append(tokoModels, model.TokoModel{
			ID:       toko.ID,
			NamaToko: toko.NamaToko,
			UrlFoto:  toko.UrlFoto,
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(pagination.Limit)))
	hasNext := pagination.Page < totalPages
	hasPrev := pagination.Page > 1

	result := &model.AllToko{
		Data:       tokoModels,
		TotalItems: total,
		TotalPages: totalPages,
		Page:       pagination.Page,
		Limit:      pagination.Limit,
		HasNext:    hasNext,
		HasPrev:    hasPrev,
	}

	return result, nil
}
