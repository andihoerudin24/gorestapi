package controller

import (
	"github.com/gin-gonic/gin"
	"gorestapi/src/apps/post/service"
	"gorestapi/utils"
	"net/http"
)

type PostController interface {
	GetAllPost(ctx *gin.Context)
}

type postController struct {
	postService service.PostService
}

func NewPostController(postRepository service.PostService) *postController {
	return &postController{postService: postRepository}
}

func (p *postController) GetAllPost(ctx *gin.Context) {
	response := utils.Response{C: ctx}
	res, err := p.postService.GetAllPost()
	if err == nil {
		response.ResponseFormatter(http.StatusOK, "list data post", nil, res)
		return
	} else {
		response.ResponseFormatter(http.StatusBadRequest, "error get data post", err.Error(), nil)
		return
	}
}
