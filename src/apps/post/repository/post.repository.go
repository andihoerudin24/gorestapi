package repository

import (
	"gorestapi/src/apps/post/model"
	"gorestapi/src/apps/post/response"
	"gorm.io/gorm"
)

type PostRepository interface {
	GetAllPost(perPage int64, Offset int64) ([]response.PostResponse, error, int64)
	CreatePost(postModel model.PostModel) (*model.PostModel, error)
	FindById(id int) (*model.PostModel, error)
}

type postRepository struct {
	connection *gorm.DB
}

func NewPostRepository(connection *gorm.DB) *postRepository {
	return &postRepository{connection: connection}
}

func (p *postRepository) GetAllPost(perPage int64, Offset int64) ([]response.PostResponse, error, int64) {
	var responsePost []response.PostResponse
	var count int64

	errs := p.connection.Debug().Table("posts").Count(&count).Error
	if errs != nil {
		return nil, errs, count
	}

	Offset = (Offset - 1) * perPage

	rows, err := p.connection.Debug().Table("posts").Select("posts.id,posts.title,posts.content,posts.slug,posts.image,users.name,users.id as user_id,users.phone").Joins("INNER JOIN users on users.id = posts.user_id").Limit(int(perPage)).Offset(int(Offset)).Rows()
	if err == nil {
		for rows.Next() {
			p.connection.ScanRows(rows, &responsePost)
		}
	}
	return responsePost, err, count
}

func (p *postRepository) CreatePost(postModel model.PostModel) (*model.PostModel, error) {
	res := p.connection.Debug().Table("posts").Create(&postModel)
	return &postModel, res.Error
}

func (p *postRepository) FindById(id int) (*model.PostModel, error) {
	var datapost model.PostModel
	res := p.connection.Debug().Table("posts").Where("id = ?", id).Find(&datapost)
	if datapost.ID != 0 {
		return &datapost, res.Error
	}
	return nil, nil

}
