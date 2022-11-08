package model

import (
	"gorestapi/utils"
)

type UserModel struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	utils.BaseModel
}

func (UserModel) TableName() string {
	return "users"
}

func NewUserModel() UserModel {
	return UserModel{}
}
