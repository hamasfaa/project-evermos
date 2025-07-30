package entity

import "time"

type Trx struct {
	ID         int       `gorm:"primaryKey;column:id;type:int;autoIncrement"`
	UserID     int       `gorm:"column:id_user;type:int;not null"`
	AlamatID   int       `gorm:"column:id_alamat;type:int;not null"`
	HargaTotal int       `gorm:"column:harga_total;type:int;not null"`
	Kode       string    `gorm:"column:kode_invoice;type:varchar(255);not null"`
	Metode     string    `gorm:"column:metode_pembayaran;type:varchar(255);not null"`
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime;autoUpdateTime"`

	// Relasi
	User      User        `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Alamat    Alamat      `gorm:"foreignKey:AlamatID;references:ID;"`
	DetailTrx []DetailTrx `gorm:"foreignKey:TrxID;references:ID;"`
}
