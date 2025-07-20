package entity

import "time"

type Kategori struct {
	ID           int       `gorm:"primaryKey;column:id;type:int;autoIncrement"`
	NamaKategori string    `gorm:"column:nama_kategori;type:varchar(255);not null"`
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime;autoUpdateTime"`

	// Relasi
	Produks []Produk `gorm:"foreignKey:KategoriID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
