package service

import (
	"context"

	"github.com/hamasfaa/project-evermos/model"
)

type TokoService interface {
	GetMyToko(ctx context.Context, userID int) (*model.MyToko, error)
}
