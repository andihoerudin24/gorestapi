package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorestapi/src/apps/post/model"
	postResponse "gorestapi/src/apps/post/response"
	"gorestapi/src/apps/post/service"
	postValidation "gorestapi/src/apps/post/validation"
	"gorestapi/utils"
	validator2 "gorestapi/validator"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

const UPLOADDIR = "postimage"
const URLSTATIC = "/api/v1/post"

type PostController interface {
	GetAllPost(ctx *gin.Context)
	CreatePost(ctx *gin.Context)
	FindById(ctx *gin.Context)
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
		responseError := validator2.BindErrors(err)
		response.ResponseFormatter(http.StatusBadRequest, "Invalid Form", responseError, nil)
		return
	}

	newFile, errorFile := upload(ctx)

	if errorFile != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "Invalid Image", gin.H{
			"image": fmt.Sprintf("%s", errorFile),
		}, nil)
		return
	}

	var title, _ = regexp.Compile(`[^a-z A-Z/0-9]`)
	resultTitle := title.ReplaceAllString(postValidate.Title, "")
	dataPost := model.NewPostModel()
	dataPost.Title = resultTitle
	dataPost.Content = postValidate.Content
	dataPost.Slug = postValidate.Slug
	dataPost.Image = fmt.Sprintf("%v", newFile)
	dataPost.UserID = uint(postValidate.UserId)
	resultInsert, err := p.postService.CreatePost(dataPost)

	var images string
	if resultInsert.Image != "" {
		images = os.Getenv("APP_HTTP") + os.Getenv("APP_URL") + ":" + os.Getenv("APP_PORT") + URLSTATIC + "/" + os.Getenv("UPLOADDIR") + "/" + UPLOADDIR + "/" + resultInsert.Image
	} else {
		images = ""
	}

	if err != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "Failed Save Data", err.Error(), nil)
		return
	}

	postResponse := postResponse.NewPostResponse()
	postResponse.ID = resultInsert.ID
	postResponse.Title = resultInsert.Title
	postResponse.Content = resultInsert.Content
	postResponse.Slug = resultInsert.Slug
	postResponse.Image = images
	postResponse.UserId = int(resultInsert.UserID)
	response.ResponseFormatter(http.StatusOK, "success save data", nil, postResponse)
}

func (p *postController) FindById(ctx *gin.Context) {
	response := utils.Response{C: ctx}
	id, errid := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if errid != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "error id", errid, gin.H{
			"errors": errid,
		})
	}
	responseData, _ := p.postService.FindById(int(id))
	if responseData == nil {
		response.ResponseFormatter(http.StatusInternalServerError, "data not found", "data not found", nil)
		return
	}

	var images string

	if responseData.Image != "" {
		images = os.Getenv("APP_HTTP") + os.Getenv("APP_URL") + ":" + os.Getenv("APP_PORT") + URLSTATIC + "/" + os.Getenv("UPLOADDIR") + "/" + UPLOADDIR + "/" + responseData.Image
	} else {
		images = ""
	}

	postResponse := postResponse.NewPostResponse()
	postResponse.ID = responseData.ID
	postResponse.Title = responseData.Title
	postResponse.Image = images
	postResponse.Content = responseData.Content
	postResponse.Slug = responseData.Slug
	postResponse.UserId = int(responseData.UserID)
	response.ResponseFormatter(http.StatusOK, "Post By Id", nil, postResponse)

}

func upload(ctx *gin.Context) (interface{}, error) {
	var newFilename string
	file, _ := ctx.FormFile("image")
	if file != nil {
		acceptImage := utils.AcceptImage(file.Header.Get("Content-Type"))
		if acceptImage != nil {
			return nil, errors.New(fmt.Sprintf("%s", acceptImage))
		} else {
			newFilename = uuid.New().String() + filepath.Ext(file.Filename)
			path := os.Getenv("UPLOADDIR") + "/" + UPLOADDIR + "/" + newFilename
			if err := ctx.SaveUploadedFile(file, path); err != nil {
				ctx.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			}

		}
	}
	return newFilename, nil
}
