package service

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
	"github.com/hamasfaa/project-evermos/model"
)

type UserService interface {
	RegisterUser(ctx context.Context, user model.RegisterModel) error
	LoginUser(ctx context.Context, phone string, password string) (*entity.User, error)
	Me(ctx context.Context, noTelp string) (*model.RegisterModel, error)
}
