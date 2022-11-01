package model

import (
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func NewUserModel() UserModel {
	return UserModel{}
}
