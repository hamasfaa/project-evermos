package service

import (
	"context"

	"github.com/hamasfaa/project-evermos/model"
)

type UserService interface {
	RegisterUser(ctx context.Context, user model.RegisterModel) error
}
