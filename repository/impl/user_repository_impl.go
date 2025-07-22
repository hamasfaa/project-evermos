package impl

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/repository"
	"gorm.io/gorm"
)

func NewUserRepositoryImpl(DB *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{DB: DB}
}

type userRepositoryImpl struct {
	DB *gorm.DB
}

func (userRepository *userRepositoryImpl) Create(ctx context.Context, user *entity.User) error {
	err := userRepository.DB.WithContext(ctx).Create(user).Error

	if err != nil {
		return err
	}
	return nil
}

func (userRepository *userRepositoryImpl) GetByPhone(ctx context.Context, phone string, password string) (*entity.User, error) {
	var user entity.User
	err := userRepository.DB.WithContext(ctx).Where("notelp = ? AND kata_sandi = ?", phone, password).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepository *userRepositoryImpl) Me(ctx context.Context, noTelp string) (*entity.User, error) {
	var user entity.User
	err := userRepository.DB.WithContext(ctx).Where("notelp = ?", noTelp).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepository *userRepositoryImpl) Update(ctx context.Context, noTelp string, user *entity.User) error {
	err := userRepository.DB.WithContext(ctx).Where("notelp = ?", noTelp).Updates(user).Error

	if err != nil {
		return err
	}
	return nil
}
