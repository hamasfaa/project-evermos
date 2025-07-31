package model

type FilterModel struct {
	Page  int    `query:"page"`
	Limit int    `query:"limit"`
	Nama  string `query:"nama"`
}

type FilterProdukModel struct {
	Page       int    `query:"page"`
	Limit      int    `query:"limit"`
	Nama       string `query:"nama_produk"`
	KategoriID int    `query:"category_id"`
	TokoID     int    `query:"toko_id"`
	MaxHarga   string `query:"max_harga"`
	MinHarga   string `query:"min_harga"`
}

type FilterTrxModel struct {
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
	Search string `query:"search"`
}
