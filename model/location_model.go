package model

type ProvinceModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CityModel struct {
	ID         string `json:"id"`
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}
