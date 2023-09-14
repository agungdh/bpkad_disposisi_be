package models

import "time"

type User struct {
	ID        uint
	Username  string `gorm:"type:varchar(255);uniqueIndex"`
	Password  string `gorm:"type:varchar(255)"`
	Role      string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
