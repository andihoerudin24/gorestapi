package controller

import (
	"github.com/gin-gonic/gin"
	"gorestapi/src/apps/post/service"
	"gorestapi/utils"
	"net/http"
	"os"
	"strconv"
)

type PostController interface {
	GetAllPost(ctx *gin.Context)
	CreatePost(ctx *gin.Context)
}

type postController struct {
	postService service.PostService
}

func NewPostController(postRepository service.PostService) *postController {
	return &postController{postService: postRepository}
}

func (p *postController) GetAllPost(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	if page <= 0 {
		page = 1
	}
	perPage, _ := strconv.Atoi(os.Getenv("PAGE"))

	response := utils.Response{C: ctx}

	res, err, totalRows := p.postService.GetAllPost(int64(perPage), int64(page))
	pagination, _ := utils.GetPaginationLinks(utils.PaginationParams{
		Path:        "post/all",
		TotalRows:   totalRows,
		PerPage:     int64(perPage),
		CurrentPage: int64(page),
	})

	if err == nil {
		response.ResponseFormatter(http.StatusOK, "list data post", nil, gin.H{
			"data":       res,
			"pagination": pagination,
		})
		return
	} else {
		response.ResponseFormatter(http.StatusBadRequest, "error get data post", err.Error(), nil)
		return
	}
}

func (p *postController) CreatePost(ctx *gin.Context) {
	response := utils.Response{C: ctx}
	response.ResponseFormatter(http.StatusOK, "siap", nil, nil)
}
