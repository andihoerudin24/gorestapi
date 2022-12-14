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
	Update(id int, post model.PostModel) (int64, *response.PostResponse)
	Delete(id int) error
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

	rows, err := p.connection.Debug().Table("posts").Select("posts.id,posts.title,posts.content,posts.slug,posts.image,users.name,users.id as user_id,users.phone").Joins("INNER JOIN users on users.id = posts.user_id").Where("users.deleted_at is null").Where("posts.deleted_at is null").Limit(int(perPage)).Offset(int(Offset)).Rows()
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

func (p *postRepository) Update(id int, post model.PostModel) (int64, *response.PostResponse) {
	var image string
	if post.Image != "" {
		image = "image"
	}
	res := p.connection.Debug().Model(&post).Select("title", "content", "slug", image, "user_id").Where("id = ? AND deleted_at is null", id).Updates(map[string]interface{}{
		"title":   post.Title,
		"content": post.Content,
		"slug":    post.Slug,
		"image":   post.Image,
		"user_id": post.UserID,
	})
	responsePost := response.NewPostResponse()
	responsePost.ID = post.ID
	responsePost.Title = post.Title
	responsePost.Content = post.Content
	responsePost.Slug = post.Slug
	responsePost.Image = post.Image
	responsePost.UserId = int(post.UserID)
	return res.RowsAffected, &responsePost
}

func (p *postRepository) Delete(id int) error {
	postModel := model.NewPostModel()
	postModel.ID = id

	if err := p.connection.Debug().Where("id = ?", id).First(&postModel).Error; err != nil {
		return err
	}
	responseDelete := p.connection.Debug().Delete(&postModel)
	return responseDelete.Error

}
