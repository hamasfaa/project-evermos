package model

type DetailTrx struct {
	ProductID int `json:"product_id"`
	Kuantitas int `json:"kuantitas"`
}

type Transaksi struct {
	Metode      string      `json:"method_bayar"`
	AlamatKirim int         `json:"alamat_kirim"`
	DetailTrx   []DetailTrx `json:"detail_trx"`
}

type Trx struct {
	ID         int            `json:"id"`
	HargaTotal int            `json:"harga_total"`
	Kode       string         `json:"kode_invoice"`
	Metode     string         `json:"method_bayar"`
	Alamat     AlamatResponse `json:"alamat_kirim"`
	Detail     []TrxDetail    `json:"detail_trx"`
}

type TrxDetail struct {
	Produk     Produk `json:"produk"`
	Kuantitas  int    `json:"kuantitas"`
	HargaTotal string `json:"harga_total"`
}

type AllTrx struct {
	Data       interface{} `json:"data"`
	TotalItems int64       `json:"total_items"`
	TotalPages int         `json:"total_pages"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	HasNext    bool        `json:"has_next"`
	HasPrev    bool        `json:"has_prev"`
}
