package service

import (
	"context"

	"github.com/hamasfaa/project-evermos/entity"
)

type AlamatService interface {
	Create(ctx context.Context, userID int, alamatData *entity.Alamat) error
}
