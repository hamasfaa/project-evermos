package impl

import (
	"context"

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
