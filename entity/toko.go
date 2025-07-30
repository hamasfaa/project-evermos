package entity

import "time"

type Toko struct {
	ID        int       `gorm:"primaryKey;column:id;type:int;autoIncrement"`
	UserID    int       `gorm:"column:id_user;type:int;not null"`
	NamaToko  string    `gorm:"column:nama_toko;type:varchar(255);not null"`
	UrlFoto   string    `gorm:"column:url_foto;type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;autoUpdateTime"`

	// Relasi
	User    User     `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Produks []Produk `gorm:"foreignKey:TokoID;references:ID;"`
	// DetailTrxs []Trx       `gorm:"foreignKey:UserID;references:ID;"`
	LogProduks []LogProduk `gorm:"foreignKey:TokoID;references:ID;"`
}
