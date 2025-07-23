package repository

import (
	"context"

	"github.com/hamasfaa/project-evermos/model"
)

type TokoRepository interface {
	Create(ctx context.Context, userID int, tokoData *model.CreateToko) error
}
