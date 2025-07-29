package model

type CreateProduct struct {
	Name          string
	KategoriID    int
	HargaReseller string
	HargaKonsumen string
	Stok          int
	Deskripsi     string
	Url           []string
	Slug          string
}

type FotoProduk struct {
	Url string
}

type AllProduk struct {
	Data       interface{} `json:"data"`
	TotalItems int64       `json:"total_items"`
	TotalPages int         `json:"total_pages"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	HasNext    bool        `json:"has_next"`
	HasPrev    bool        `json:"has_prev"`
}

type Produk struct {
	ID            int                  `json:"id"`
	NamaProduk    string               `json:"nama_produk"`
	Slug          string               `json:"slug"`
	HargaReseller string               `json:"harga_reseller"`
	HargaKonsumen string               `json:"harga_konsumen"`
	Stok          int                  `json:"stok"`
	Deskripsi     string               `json:"deskripsi"`
	Toko          TokoModel            `json:"toko"`
	Kategori      KategoriResponse     `json:"kategori"`
	Foto          []FotoProdukResponse `json:"photos"`
}

type FotoProdukResponse struct {
	ID       int    `json:"id"`
	ProdukID int    `json:"produk_id"`
	Url      string `json:"url"`
}
