package migration

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Slug      string `json:"slug"`
	Image     string `json:"image"`
	UserId    uint   `json:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
