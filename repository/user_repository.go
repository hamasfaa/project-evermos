package repository

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
}
