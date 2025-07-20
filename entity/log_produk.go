package entity

import "time"

type LogProduk struct {
	ID            int       `gorm:"primaryKey;column:id;type:int;autoIncrement"`
	ProdukID      int       `gorm:"column:id_produk;type:int;not null"`
	NamaProduk    string    `gorm:"column:nama_produk;type:varchar(255);not null"`
	Slug          string    `gorm:"column:slug;type:varchar(255);not null"`
	HargaReseller string    `gorm:"column:harga_reseller;type:varchar(255);not null"`
	HargaKonsumen string    `gorm:"column:harga_konsumen;type:varchar(255);not null"`
	Deskripsi     string    `gorm:"column:deskripsi;type:text;not null"`
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime;autoCreateTime"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime;autoUpdateTime"`
	TokoID        int       `gorm:"column:id_toko;type:int;not null"`
	KategoriID    int       `gorm:"column:id_kategori;type:int;not null"`

	// Relasi
	Produk    Produk    `gorm:"foreignKey:ProdukID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	DetailTrx DetailTrx `gorm:"foreignKey:LogID;references:ID;"`
}
