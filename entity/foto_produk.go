package entity

import "time"

type FotoProduk struct {
	ID        int       `gorm:"primaryKey;column:id;type:int;autoIncrement"`
	ProdukID  int       `gorm:"column:id_produk;type:int;not null"`
	Url       string    `gorm:"column:url;type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;autoUpdateTime"`

	// Relasi
	Produk Produk `gorm:"foreignKey:ProdukID;references:ID;"`
}
