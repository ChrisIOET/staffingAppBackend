package model

import (
	"time"
)
type Category struct {
	ID        uint      `gorm:"serial primary key" json:"id"`
	Name      string    `gorm:"uniqueIndex:udx_name" json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}
