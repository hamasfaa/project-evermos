package service

import (
	"context"

	"github.com/hamasfaa/project-evermos/model"
)

type LocationService interface {
	GetProvinceByID(ctx context.Context, provinceID string) (*model.ProvinceModel, error)
	GetCityByID(ctx context.Context, provinceID, cityID string) (*model.CityModel, error)
	GetProvince(ctx context.Context) ([]model.ProvinceModel, error)
	GetCities(ctx context.Context, provinceID string) ([]model.CityModel, error)
}
