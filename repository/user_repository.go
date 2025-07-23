package repository

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (int, error)
	GetByPhone(ctx context.Context, phone string, password string) (*entity.User, error)
	Me(ctx context.Context, noTelp string) (*entity.User, error)
	Update(ctx context.Context, noTelp string, user *entity.User) error
}
