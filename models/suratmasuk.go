package models

import "time"

type SuratMasuk struct {
	ID               uint
	NomorSurat       string `gorm:"type:varchar(255)"`
	TanggalSurat     time.Time
	InstansiPengirim string `gorm:"type:varchar(255)"`
	Perihal          string `gorm:"type:varchar(255)"`
	PejabatPengirim  string `gorm:"type:varchar(255)"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
