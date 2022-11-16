package controller

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorestapi/src/apps/user/model"
	"gorestapi/src/apps/user/service"
	"gorestapi/src/apps/user/validation"
	"gorestapi/utils"
	validator2 "gorestapi/validator"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

const UPLOADDIR = "user"
const URLSTATIC = "/api/v1/user"

type UserController interface {
	GetAllUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type userController struct {
	services service.UserService
}

func NewUserController(services service.UserService) *userController {
	return &userController{services: services}
}

func (u *userController) GetAllUser(ctx *gin.Context) {
	response := utils.Response{C: ctx}
	var pagination interface{}
	page, _ := strconv.Atoi(ctx.Query("page"))
	if page <= 0 {
		page = 1
	}
	perPage := 5
	DataUser, totalrows := u.services.GetAllUser(int64(perPage), int64(page))

	pagination, _ = utils.GetPaginationLinks(utils.PaginationParams{
		Path:        "user/all",
		TotalRows:   totalrows,
		PerPage:     int64(perPage),
		CurrentPage: int64(page),
	})

	if len(*DataUser) == 0 {
		response.ResponseFormatter(http.StatusNotFound, "data not found", DataUser, nil)
		return
	}

	var dataresponse []interface{}
	var images string

	for _, value := range *DataUser {
		if value.Image != "" {
			images = os.Getenv("APP_HTTP") + os.Getenv("APP_URL") + ":" + os.Getenv("APP_PORT") + URLSTATIC + "/" + os.Getenv("UPLOADDIR") + "/" + UPLOADDIR + "/" + value.Image
		} else {
			images = ""
		}
		dataresponse = append(dataresponse, map[string]interface{}{
			"id":      value.ID,
			"name":    value.Name,
			"phone":   value.Phone,
			"email":   value.Email,
			"address": value.Address,
			"image":   images,
		})
	}
	response.ResponseFormatter(http.StatusOK, "List User", nil, gin.H{
		"data":       dataresponse,
		"pagination": pagination,
	})
}

func (u *userController) CreateUser(ctx *gin.Context) {
	response := utils.Response{C: ctx}
	Uvalidation := validation.NewCreateUserValidation()
	newFileName, errorFilename := upload(ctx)

	if errorFilename != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "invalid image", gin.H{
			"image": fmt.Sprintf("%s", errorFilename),
		}, nil)
		return
	}

	if err := Uvalidation.Bind(ctx); err != nil {
		fmt.Println("errr", err)
		ResponError := validator2.BindErrors(err)
		response.ResponseFormatter(http.StatusBadRequest, "Invalid Form", ResponError, nil)
		return
	}

	newUsr := model.NewUserModel()
	newUsr.Name = Uvalidation.Name
	newUsr.Email = Uvalidation.Email
	newUsr.Phone = Uvalidation.Phone
	newUsr.Address = Uvalidation.Address
	newUsr.Image = fmt.Sprintf("%v", newFileName)
	res, err := u.services.CreateUser(newUsr)

	if err != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "Failed Save Data", err.Error(), gin.H{
			"errors": err.Error(),
		})
		return
	}

	response.ResponseFormatter(http.StatusOK, "Success Save Data", nil, gin.H{
		"email":   res.Email,
		"name":    res.Name,
		"phone":   res.Phone,
		"address": res.Address,
		"image":   res.Image,
	})
}

func (u *userController) FindById(ctx *gin.Context) {
	response := utils.Response{C: ctx}
	var images string
	id, errid := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if errid != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "error id", errid, gin.H{
			"errors": errid,
		})
	}
	responses, _ := u.services.FindById(id)
	if responses.Image != "" {
		images = os.Getenv("APP_HTTP") + os.Getenv("APP_URL") + ":" + os.Getenv("APP_PORT") + URLSTATIC + "/" + os.Getenv("UPLOADDIR") + "/" + UPLOADDIR + "/" + responses.Image
	} else {
		images = ""
	}
	dataresponse := map[string]interface{}{
		"id":      responses.ID,
		"email":   responses.Email,
		"name":    responses.Name,
		"phone":   responses.Phone,
		"address": responses.Address,
		"image":   images,
	}
	if responses == nil {
		response.ResponseFormatter(http.StatusInternalServerError, "data not found", "data not found", nil)
		return
	}

	response.ResponseFormatter(http.StatusOK, "User By Id", nil, dataresponse)
}

func (u *userController) Update(ctx *gin.Context) {
	response := utils.Response{C: ctx}
	id, errid := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if errid != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "error id", errid, gin.H{
			"errors": errid,
		})
	}

	responses, _ := u.services.FindById(id)
	if responses == nil {
		response.ResponseFormatter(http.StatusInternalServerError, "data not found", "data not found", nil)
		return
	}

	newfilename, errorfilename := upload(ctx)
	if errorfilename != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "Invalid Image", gin.H{
			"image": fmt.Sprintf("%s", errorfilename),
		}, nil)
	}

	Uvalidation := validation.NewUpdateUserValidation()
	if err := Uvalidation.Bind(ctx); err != nil {
		ResponError := validator2.BindErrors(err)

		response.ResponseFormatter(http.StatusBadRequest, "Invalid Form", ResponError, nil)
		return
	}
	newUsr := model.NewUserModel()
	newUsr.ID = uint(id)
	newUsr.Name = Uvalidation.Name
	newUsr.Email = Uvalidation.Email
	newUsr.Phone = Uvalidation.Phone
	newUsr.Address = Uvalidation.Address
	newUsr.Image = fmt.Sprintf("%v", newfilename)
	_ = u.services.Update(id, newUsr)

	updData := map[string]interface{}{}
	v := reflect.ValueOf(newUsr)
	typeOfV := v.Type()
	for i := 0; i < v.NumField(); i++ {
		updData[strings.ToLower(typeOfV.Field(i).Name)] = v.Field(i).Interface()
	}
	delete(updData, "basemodel")
	response.ResponseFormatter(http.StatusOK, fmt.Sprintf("sukses update data with id = %s", strconv.Itoa(int(id))), nil, updData)
}

func upload(ctx *gin.Context) (interface{}, error) {
	var newFileName string
	file, _ := ctx.FormFile("image")
	if file != nil {
		acceptImage := utils.AcceptImage(file.Header.Get("Content-Type"))
		if acceptImage != nil {
			return nil, errors.New(fmt.Sprintf("%s", acceptImage))
		} else {
			newFileName = uuid.New().String() + filepath.Ext(file.Filename)
			path := os.Getenv("UPLOADDIR") + "/" + UPLOADDIR + "/" + newFileName
			if err := ctx.SaveUploadedFile(file, path); err != nil {
				ctx.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			}
		}
	}
	return newFileName, nil
}

func (u *userController) Delete(ctx *gin.Context) {
	response := utils.Response{C: ctx}
	IdUser, errid := strconv.Atoi(ctx.Param("id"))
	if errid != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "error id", errid, gin.H{
			"errors": errid,
		})
	}
	data := u.services.Delete(int64(IdUser))
	if data != nil {
		response.ResponseFormatter(http.StatusNotFound, fmt.Sprintf("%v", data), nil, nil)
		return
	}
	response.ResponseFormatter(http.StatusOK, "Success Delete Data", nil, nil)
}
