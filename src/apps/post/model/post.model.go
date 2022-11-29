package model

import (
	"gorestapi/src/apps/user/model"
	"gorestapi/utils"
)

type PostModel struct {
	utils.BaseModel
	Title     string          `json:"title"`
	Content   string          `json:"content"`
	Slug      string          `json:"slug"`
	Image     string          `json:"image"`
	UserID    uint            `json:"user_id"`
	UserModel model.UserModel `gorm:"foreignKey:UserID"`
}

func (PostModel) TableName() string {
	return "posts"
}

func NewPostModel() PostModel {
	return PostModel{}
}
