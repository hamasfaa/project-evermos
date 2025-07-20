package entity

import "time"

type Alamat struct {
	ID           int       `gorm:"primaryKey;column:id;type:int;autoIncrement"`
	UserID       int       `gorm:"column:id_user;type:int;not null"`
	JudulAlamat  string    `gorm:"column:judul_alamat;type:varchar(255);not null"`
	NamaPenerima string    `gorm:"column:nama_penerima;type:varchar(255);not null"`
	Notelp       string    `gorm:"column:notelp;type:varchar(20);not null"`
	DetailAlamat string    `gorm:"column:detail_alamat;type:text;not null"`
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime;autoUpdateTime"`

	// Relasi
	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Trx  *Trx `gorm:"foreignKey:AlamatID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
