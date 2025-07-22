package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/hamasfaa/project-evermos/model"
	"github.com/hamasfaa/project-evermos/service"
)

type locationServiceImpl struct {
	httpClient *http.Client
}

func NewLocationServiceImpl() service.LocationService {
	return &locationServiceImpl{
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (s *locationServiceImpl) GetProvinceByID(ctx context.Context, provinceID string) (*model.ProvinceModel, error) {
	resp, err := s.httpClient.Get("https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var provinces []model.ProvinceModel
	if err := json.NewDecoder(resp.Body).Decode(&provinces); err != nil {
		return nil, err
	}

	for _, province := range provinces {
		if province.ID == provinceID {
			return &province, nil
		}
	}

	return nil, fmt.Errorf("province not found")
}

func (s *locationServiceImpl) GetCityByID(ctx context.Context, provinceID, cityID string) (*model.CityModel, error) {
	url := fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/regencies/%s.json", provinceID)
	resp, err := s.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cities []model.CityModel
	if err := json.NewDecoder(resp.Body).Decode(&cities); err != nil {
		return nil, err
	}

	for _, city := range cities {
		if city.ID == cityID {
			return &city, nil
		}
	}

	return nil, fmt.Errorf("city not found")
}

func (s *locationServiceImpl) GetProvince(ctx context.Context) ([]model.ProvinceModel, error) {
	resp, err := s.httpClient.Get("https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var provinces []model.ProvinceModel
	if err := json.NewDecoder(resp.Body).Decode(&provinces); err != nil {
		return nil, err
	}

	return provinces, nil
}

func (s *locationServiceImpl) GetCities(ctx context.Context, provinceID string) ([]model.CityModel, error) {
	url := fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/regencies/%s.json", provinceID)
	resp, err := s.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cities []model.CityModel
	if err := json.NewDecoder(resp.Body).Decode(&cities); err != nil {
		return nil, err
	}

	return cities, nil
}
