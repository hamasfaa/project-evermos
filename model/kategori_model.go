package model

type Kategori struct {
	NamaKategori string `json:"nama_category"`
}

type KategoriResponse struct {
	ID           int    `json:"id"`
	NamaKategori string `json:"nama_category"`
}
