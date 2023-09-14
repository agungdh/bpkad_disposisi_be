package models

import "time"

type Bidang struct {
	ID        uint
	Bidang    string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
