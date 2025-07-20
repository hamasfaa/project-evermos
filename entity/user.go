package entity

import "time"

type User struct {
	ID           int       `gorm:"primaryKey;column:id;type:int;autoIncrement"`
	Nama         string    `gorm:"column:nama;type:varchar(255);not null"`
	Kata_Sandi   string    `gorm:"column:kata_sandi;type:varchar(255);not null"`
	Notelp       string    `gorm:"column:notelp;type:varchar(20);unique;not null"`
	TanggalLahir string    `gorm:"column:tanggal_lahir;type:date"`
	JenisKelamin string    `gorm:"column:jenis_kelamin;type:varchar(255)"`
	Tentang      string    `gorm:"column:tentang;type:text"`
	Pekerjaan    string    `gorm:"column:pekerjaan;type:varchar(255)"`
	Email        string    `gorm:"column:email;type:varchar(255);unique"`
	IDProvinsi   string    `gorm:"column:id_provinsi;type:varchar(255)"`
	IDKota       string    `gorm:"column:id_kota;type:varchar(255)"`
	IsAdmin      bool      `gorm:"column:isAdmin;type:bool;default:false"`
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime;autoUpdateTime"`

	// Relasi
	Alamats []Alamat `gorm:"foreignKey:UserID;references:ID"`
	Toko    *Toko    `gorm:"foreignKey:UserID;references:ID"`
	Trx     *Trx     `gorm:"foreignKey:UserID;references:ID"`
}
