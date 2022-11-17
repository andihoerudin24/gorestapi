package repository

import (
	"fmt"
	"gorestapi/src/apps/post/model"
	"gorm.io/gorm"
)

type PostRepository interface {
	GetAllPost() ([]model.PostModel, error)
}

type postRepository struct {
	connection *gorm.DB
}

func NewPostRepository(connection *gorm.DB) *postRepository {
	return &postRepository{connection: connection}
}

func (p *postRepository) GetAllPost() ([]model.PostModel, error) {
	prop := model.PostModel{}
	if err := p.connection.Debug().Model(&prop).Joins("UserModel").Find(&prop).Error; err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", prop)
	return nil, nil
}
