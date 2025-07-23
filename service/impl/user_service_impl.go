package impl

import (
	"context"
	"errors"
	"time"

	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/model"
	"github.com/hamasfaa/project-evermos/repository"
	"github.com/hamasfaa/project-evermos/service"
)

var ErrInvalidDateFormat = errors.New("invalid date format, expected DD/MM/YYYY")

func NewUserServiceImpl(userRepository *repository.UserRepository, tokoRepository *repository.TokoRepository) service.UserService {
	return &userServiceImpl{UserRepository: *userRepository, TokoRepository: *tokoRepository}
}

type userServiceImpl struct {
	repository.UserRepository
	repository.TokoRepository
}

func (userService *userServiceImpl) RegisterUser(ctx context.Context, user model.RegisterModel) error {
	parseDate, err := time.Parse("02/01/2006", user.TanggalLahir)
	if err != nil {
		return ErrInvalidDateFormat
	}
	userEntity := entity.User{
		Nama:         user.Nama,
		KataSandi:    user.KataSandi,
		Notelp:       user.NoTelp,
		TanggalLahir: parseDate,
		Pekerjaan:    user.Pekerjaan,
		Email:        user.Email,
		IDProvinsi:   user.IDProvinsi,
		IDKota:       user.IDKota,
	}

	userID, err := userService.UserRepository.Create(ctx, &userEntity)
	if err != nil {
		return err
	}

	tokoModel := model.CreateToko{
		UserID:   userID,
		NamaToko: user.Nama,
		UrlFoto:  "",
	}

	err = userService.TokoRepository.Create(ctx, userID, &tokoModel)
	if err != nil {
		return err
	}

	return nil
}

func (userService *userServiceImpl) LoginUser(ctx context.Context, phone string, password string) (*entity.User, error) {
	userEntity, err := userService.UserRepository.GetByPhone(ctx, phone, password)
	if err != nil {
		return nil, err
	}
	return userEntity, nil
}

func (userService *userServiceImpl) Me(ctx context.Context, noTelp string) (*model.RegisterModel, error) {
	userEntity, err := userService.UserRepository.Me(ctx, noTelp)
	if err != nil {
		return nil, err
	}

	meResponse := &model.RegisterModel{
		Nama:         userEntity.Nama,
		KataSandi:    userEntity.KataSandi,
		NoTelp:       userEntity.Notelp,
		TanggalLahir: userEntity.TanggalLahir.Format("02/01/2006"),
		Pekerjaan:    userEntity.Pekerjaan,
		Email:        userEntity.Email,
		IDProvinsi:   userEntity.IDProvinsi,
		IDKota:       userEntity.IDKota,
	}

	return meResponse, nil
}

func (userService *userServiceImpl) UpdateUser(ctx context.Context, noTelp string, user model.RegisterModel) error {
	parseDate, err := time.Parse("02/01/2006", user.TanggalLahir)
	if err != nil {
		return ErrInvalidDateFormat
	}

	userEntity := entity.User{
		Nama:         user.Nama,
		Notelp:       user.NoTelp,
		TanggalLahir: parseDate,
		Pekerjaan:    user.Pekerjaan,
		Email:        user.Email,
		IDProvinsi:   user.IDProvinsi,
		IDKota:       user.IDKota,
	}

	err = userService.UserRepository.Update(ctx, noTelp, &userEntity)
	if err != nil {
		return err
	}

	return nil
}
