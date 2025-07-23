package service

import (
	"context"

	"github.com/hamasfaa/project-evermos/model"
)

type TokoService interface {
	GetMyToko(ctx context.Context, userID int) (*model.MyToko, error)
	GetTokoByID(ctx context.Context, tokoID int) (*model.TokoModel, error)
	GetAllTokos(ctx context.Context, pagination model.FilterModel) (*model.AllToko, error)
	UpdateToko(ctx context.Context, tokoData model.CreateToko, tokoID int) error
}
