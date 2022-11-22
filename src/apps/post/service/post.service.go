package service

import (
	"gorestapi/src/apps/post/repository"
	"gorestapi/src/apps/post/response"
)

type PostService interface {
	GetAllPost() ([]response.PostResponse, error)
}

type postService struct {
	repository repository.PostRepository
}

func NewPostService(repository repository.PostRepository) *postService {
	return &postService{repository: repository}
}

func (p postService) GetAllPost() ([]response.PostResponse, error) {
	postResponse, error := p.repository.GetAllPost()
	return postResponse, error
}
