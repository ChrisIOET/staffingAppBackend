package model

import (
	"gorm.io/gorm"
	"time"
)
type User struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Username     string         `gorm:"unique;index;not null" json:"username"`
	Password     string         `gorm:"type:varchar(1000);not null" json:"password"`
	EnglishLevel string         `json:"englishlevel"`
	Roles        string         `json:"roles"`
	IsActive     bool           `json:"isactive"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
