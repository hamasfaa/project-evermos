package model

type LoginModel struct {
	Notelp    string `json:"no_telp"`
	KataSandi string `json:"kata_sandi"`
}

type RegisterModel struct {
	Nama         string `json:"nama"`
	KataSandi    string `json:"kata_sandi"`
	NoTelp       string `json:"no_telp"`
	TanggalLahir string `json:"tanggal_lahir"`
	Pekerjaan    string `json:"pekerjaan"`
	Email        string `json:"email"`
	IDProvinsi   string `json:"id_provinsi"`
	IDKota       string `json:"id_kota"`
}

type UserResponse struct {
	Nama         string         `json:"nama"`
	NoTelp       string         `json:"no_telp"`
	TanggalLahir string         `json:"tanggal_lahir"`
	Tentang      string         `json:"tentang"`
	Pekerjaan    string         `json:"pekerjaan"`
	Email        string         `json:"email"`
	IDProvinsi   *ProvinceModel `json:"id_provinsi"`
	IDKota       *CityModel     `json:"id_kota"`
	Token        string         `json:"token"`
}
