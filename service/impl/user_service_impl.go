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

func NewUserServiceImpl(userRepository *repository.UserRepository) service.UserService {
	return &userServiceImpl{UserRepository: *userRepository}
}

type userServiceImpl struct {
	repository.UserRepository
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

	err = userService.UserRepository.Create(ctx, &userEntity)
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
