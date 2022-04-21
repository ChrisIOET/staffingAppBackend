package model

import (
	"gorm.io/gorm"
	"time"
)

// User struct
type UserSkill struct {
	UserID     uint           `gorm:"primaryKey" json:"id"`
	SkillID    uint           `gorm:"primaryKey" json:"skillid"`
	Level      string         `json:"level"`
	Detail     string         `json:"detail"`
	LastUpdate time.Time      `json:"Lastupdate"`
	VerifiedBy string         `json:"verifiedby"`
	Status     string         `json:"status"`
	CreatedAt  time.Time      `json:"createdAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
