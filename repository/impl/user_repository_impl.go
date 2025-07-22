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
