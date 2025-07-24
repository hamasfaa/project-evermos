package impl

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/model"
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

func (service *alamatServiceImpl) GetAlamatByUserID(ctx context.Context, userID int) ([]model.AlamatResponse, error) {
	alamatList, err := service.AlamatRepository.GetAlamatByUserID(ctx, userID)

	var alamatResponse []model.AlamatResponse

	for _, alamat := range alamatList {
		alamatResponse = append(alamatResponse, model.AlamatResponse{
			ID:           alamat.ID,
			JudulAlamat:  alamat.JudulAlamat,
			NamaPenerima: alamat.NamaPenerima,
			NoTelp:       alamat.Notelp,
			DetailAlamat: alamat.DetailAlamat,
		})
	}

	if err != nil {
		return nil, err
	}
	return alamatResponse, nil
}

func (service *alamatServiceImpl) GetAlamatByID(ctx context.Context, id int, userID int) (*model.AlamatResponse, error) {
	alamat, err := service.AlamatRepository.GetAlamatByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	alamatResponse := &model.AlamatResponse{
		ID:           alamat.ID,
		JudulAlamat:  alamat.JudulAlamat,
		NamaPenerima: alamat.NamaPenerima,
		NoTelp:       alamat.Notelp,
		DetailAlamat: alamat.DetailAlamat,
	}

	return alamatResponse, nil
}

func (service *alamatServiceImpl) DeleteAlamatByID(ctx context.Context, id int, userID int) error {
	err := service.AlamatRepository.DeleteAlamatByID(ctx, id, userID)
	if err != nil {
		return err
	}
	return nil
}

func (service *alamatServiceImpl) UpdateAlamatByID(ctx context.Context, id int, userID int, alamatData *model.UpdateAlamatModel) error {
	alamatEntity := &entity.Alamat{
		NamaPenerima: alamatData.NamaPenerima,
		Notelp:       alamatData.NoTelp,
		DetailAlamat: alamatData.DetailAlamat,
	}

	err := service.AlamatRepository.UpdateAlamatByID(ctx, id, userID, alamatEntity)
	if err != nil {
		return err
	}
	return nil
}
