package model

import (
	"github.com/gosimple/slug"
	"gorestapi/src/apps/user/model"
	"gorestapi/utils"
	"time"
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

func (m *PostModel) BeforeInsert() error {
	now := time.Now()
	if m.CreatedAt.IsZero() {
		m.CreatedAt = now
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = now
	}
	if m.Slug != "" {
		m.Slug = slug.Make(m.Slug)
	}

	return nil
}

func (m *PostModel) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	if m.Slug != "" {
		m.Slug = slug.Make(m.Slug)
	}
	return nil
}
