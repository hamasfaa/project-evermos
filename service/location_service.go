package service

import (
	"context"

	"github.com/hamasfaa/project-evermos/model"
)

type LocationService interface {
	GetProvinceByID(ctx context.Context, provinceID string) (*model.ProvinceModel, error)
	GetCityByID(ctx context.Context, provinceID, cityID string) (*model.CityModel, error)
}
