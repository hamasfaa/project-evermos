package impl

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/repository"
	"github.com/hamasfaa/project-evermos/service"
)

func NewAlamatServiceImpl(alamatRepository *repository.AlamatRepository) service.AlamatService {
	return &alamatServiceImpl{AlamatRepository: *alamatRepository}
}

type alamatServiceImpl struct {
	repository.AlamatRepository
}

func (service *alamatServiceImpl) Create(ctx context.Context, userID int, alamatData *entity.Alamat) error {
	alamatData.UserID = userID

	err := service.AlamatRepository.Create(ctx, alamatData)
	if err != nil {
		return err
	}
	return nil
}
