package models

import "time"

type Bidang struct {
	Id        int64     `gorm:"primaryKey" json:"id"`
	Bidang    string    `gorm:"type:varchar(255)" json:"bidang"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
