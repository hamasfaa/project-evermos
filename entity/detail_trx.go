package entity

import "time"

type DetailTrx struct {
	ID         int       `gorm:"primaryKey;column:id;type:int;autoIncrement"`
	TrxID      int       `gorm:"column:id_trx;type:int;not null"`
	TokoID     int       `gorm:"column:id_toko;type:int;not null"`
	Kuantitas  int       `gorm:"column:kuantitas;type:int;not null"`
	HargaTotal int       `gorm:"column:harga_total;type:int;not null"`
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime;autoUpdateTime"`

	// Relasi
	Toko Toko        `gorm:"foreignKey:TokoID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Trx  Trx         `gorm:"foreignKey:TrxID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Logs []LogProduk `gorm:"foreignKey:DetailTrxID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
