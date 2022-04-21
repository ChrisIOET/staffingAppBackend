package model

import (
	"time"
	"gorm.io/gorm"
)

// Id, Name, CategoryId

type Skill struct {
	ID		uint		`json:"id"`	
	Name		string	`json:"name"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
