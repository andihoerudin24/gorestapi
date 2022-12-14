package service

import (
	"gorestapi/src/apps/post/model"
	"gorestapi/src/apps/post/repository"
	"gorestapi/src/apps/post/response"
)

type PostService interface {
	GetAllPost(perPage int64, Page int64) ([]response.PostResponse, error, int64)
	CreatePost(postModel model.PostModel) (*model.PostModel, error)
	FindById(id int) (*model.PostModel, error)
	Update(id int, postModel model.PostModel) (int64, *response.PostResponse)
	Delete(id int) error
}

type postService struct {
	repository repository.PostRepository
}

func NewPostService(repository repository.PostRepository) *postService {
	return &postService{repository: repository}
}

func (p *postService) GetAllPost(perPage int64, Page int64) ([]response.PostResponse, error, int64) {
	postResponse, error, count := p.repository.GetAllPost(perPage, Page)
	return postResponse, error, count
}

func (p *postService) CreatePost(postModel model.PostModel) (*model.PostModel, error) {
	postResponse, err := p.repository.CreatePost(postModel)
	return postResponse, err
}

func (p *postService) FindById(id int) (*model.PostModel, error) {
	postResponse, err := p.repository.FindById(id)
	return postResponse, err
}

func (p *postService) Update(id int, postModel model.PostModel) (int64, *response.PostResponse) {
	postUpdate, postResponse := p.repository.Update(id, postModel)
	return postUpdate, postResponse
}

func (p *postService) Delete(id int) error {
	postDelete := p.repository.Delete(id)
	return postDelete
}
