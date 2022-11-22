package repository

import (
	"gorestapi/src/apps/post/response"
	"gorm.io/gorm"
)

type PostRepository interface {
	GetAllPost() ([]response.PostResponse, error)
}

type postRepository struct {
	connection *gorm.DB
}

func NewPostRepository(connection *gorm.DB) *postRepository {
	return &postRepository{connection: connection}
}

func (p *postRepository) GetAllPost() ([]response.PostResponse, error) {
	var responsePost []response.PostResponse
	rows, err := p.connection.Debug().Table("posts").Select("posts.title,posts.content,posts.slug,posts.image,users.name,users.id as user_id,users.phone").Joins("INNER JOIN users on users.id = posts.user_id").Rows()
	if err == nil {
		for rows.Next() {
			p.connection.ScanRows(rows, &responsePost)
		}
	}
	return responsePost, err
}
