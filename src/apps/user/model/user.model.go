package model

import "gorestapi/utils"

type UserModel struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Image   string `json:"image"`
	utils.BaseModel
}

func (UserModel) TableName() string {
	return "users"
}

func NewUserModel() UserModel {
	return UserModel{}
}
