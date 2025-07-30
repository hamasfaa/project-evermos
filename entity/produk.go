package entity

import "time"

type Produk struct {
	ID            int       `gorm:"primaryKey;column:id;type:int;autoIncrement"`
	NamaProduk    string    `gorm:"column:nama_produk;type:varchar(255);not null"`
	Slug          string    `gorm:"column:slug;type:varchar(255);unique;not null"`
	HargaReseller string    `gorm:"column:harga_reseller;type:varchar(255);not null"`
	HargaKonsumen string    `gorm:"column:harga_konsumen;type:varchar(255);not null"`
	Stok          int       `gorm:"column:stok;type:int;not null"`
	Deskripsi     string    `gorm:"column:deskripsi;type:text;not null"`
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime;autoCreateTime"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime;autoUpdateTime"`
	TokoID        int       `gorm:"column:id_toko;type:int;not null"`
	KategoriID    int       `gorm:"column:id_kategori;type:int;not null"`

	// Relasi
	Toko        Toko         `gorm:"foreignKey:TokoID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	LogProduk   *LogProduk   `gorm:"foreignKey:ProdukID;references:ID;"`
	FotoProduks []FotoProduk `gorm:"foreignKey:ProdukID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Kategori    Kategori     `gorm:"foreignKey:KategoriID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	DetailTrxs  []DetailTrx  `gorm:"foreignKey:ProdukID;references:ID"`
}
