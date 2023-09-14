package models

import "time"

type Pegawai struct {
	ID              uint
	Nama            string `type:varchar(255);gorm:"uniqueIndex"`
	Nip             string `gorm:"type:varchar(255)"`
	Eselon          string `gorm:"type:varchar(255)"`
	PangkatGolongan string `gorm:"type:varchar(255)"`
	JabatanID       uint
	Jabatan         Jabatan
	UserID          uint
	User            User
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
