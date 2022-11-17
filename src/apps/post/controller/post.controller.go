package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorestapi/src/apps/post/repository"
	"net/http"
)

type PostController interface {
	GetAllPost(ctx *gin.Context)
}

type postController struct {
	postrepository repository.PostRepository
}

func NewPostController(postRepository repository.PostRepository) *postController {
	return &postController{postrepository: postRepository}
}

func (p *postController) GetAllPost(ctx *gin.Context) {
	res, err := p.postrepository.GetAllPost()
	fmt.Println("rescontroller", res)
	fmt.Println("err", err)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
