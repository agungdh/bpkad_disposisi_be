package models

import "time"

type Jabatan struct {
	ID                uint
	Jabatan           string `gorm:"type:varchar(255)"`
	BidangID          uint
	Bidang            Bidang
	IsKepalaBadan     int `gorm:"type:smallint(1)"`
	IsSekretaris      int `gorm:"type:smallint(1)"`
	IsKepalaBidang    int `gorm:"type:smallint(1)"`
	IsKepalaSubBagian int `gorm:"type:smallint(1)"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
