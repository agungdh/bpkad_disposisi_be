package models

import "time"

type Honorer struct {
	ID        uint
	Nama      string `gorm:"type:varchar(255)"`
	Nik       string `gorm:"type:varchar(255);uniqueIndex"`
	JabatanID uint
	Jabatan   Jabatan
	UserID    uint
	User      User
	CreatedAt time.Time
	UpdatedAt time.Time
}
