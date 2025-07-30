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
