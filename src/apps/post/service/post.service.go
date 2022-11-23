package service

import (
	"gorestapi/src/apps/post/repository"
	"gorestapi/src/apps/post/response"
)

type PostService interface {
	GetAllPost(perPage int64, Page int64) ([]response.PostResponse, error, int64)
}

type postService struct {
	repository repository.PostRepository
}

func NewPostService(repository repository.PostRepository) *postService {
	return &postService{repository: repository}
}

func (p postService) GetAllPost(perPage int64, Page int64) ([]response.PostResponse, error, int64) {
	postResponse, error, count := p.repository.GetAllPost(perPage, Page)
	return postResponse, error, count
}
