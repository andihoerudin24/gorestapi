package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorestapi/src/apps/post/model"
	postResponse "gorestapi/src/apps/post/response"
	"gorestapi/src/apps/post/service"
	postValidation "gorestapi/src/apps/post/validation"
	"gorestapi/utils"
	validator2 "gorestapi/validator"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

type PostController interface {
	GetAllPost(ctx *gin.Context)
	CreatePost(ctx *gin.Context)
}

type postController struct {
	postService service.PostService
}

func NewPostController(postService service.PostService) *postController {
	return &postController{postService: postService}
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
	postValidate := postValidation.NewCreatePostValidation()
	if err := postValidate.Bind(ctx); err != nil {
		fmt.Println("Err", err)
		responseError := validator2.BindErrors(err)
		response.ResponseFormatter(http.StatusBadRequest, "Invalid Form", responseError, nil)
		return
	}
	var title, _ = regexp.Compile(`[^a-z A-Z/0-9]`)
	resultTitle := title.ReplaceAllString(postValidate.Title, "")
	dataPost := model.NewPostModel()
	dataPost.Title = resultTitle
	dataPost.Content = postValidate.Content
	dataPost.Slug = postValidate.Slug
	dataPost.UserID = uint(postValidate.UserId)
	resultInsert, err := p.postService.CreatePost(dataPost)

	if err != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "Failed Save Data", err.Error(), nil)
		return
	}

	postResponse := postResponse.NewPostResponse()
	postResponse.ID = resultInsert.ID
	postResponse.Title = resultInsert.Title
	postResponse.Content = resultInsert.Content
	postResponse.Slug = resultInsert.Slug
	postResponse.Image = resultInsert.Image
	postResponse.UserId = int(resultInsert.UserID)

	response.ResponseFormatter(http.StatusOK, "success save data", nil, postResponse)
}
