package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorestapi/src/apps/user/model"
	"gorestapi/src/apps/user/service"
	"gorestapi/src/apps/user/validation"
	"gorestapi/utils"
	validator2 "gorestapi/validator"
	"net/http"
	"reflect"
	"strconv"
)

type UserController interface {
	GetAllUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Update()
}

type userController struct {
	services service.UserService
}

func NewUserController(services service.UserService) *userController {
	return &userController{services: services}
}

func (u *userController) GetAllUser(ctx *gin.Context) {
	response := utils.Response{C: ctx}
	page, _ := strconv.Atoi(ctx.Query("page"))
	if page <= 0 {
		page = 1
	}
	perPage := 5
	DataUser, totalrows := u.services.GetAllUser(int64(perPage), int64(page))

	pagination, _ := utils.GetPaginationLinks(utils.PaginationParams{
		Path:        "user/all",
		TotalRows:   totalrows,
		PerPage:     int64(perPage),
		CurrentPage: int64(page),
	})

	if DataUser == nil {
		response.ResponseFormatter(http.StatusNotFound, "data not found", DataUser, nil)
	}
	response.ResponseFormatter(http.StatusOK, "List User", nil, gin.H{
		"data":       DataUser,
		"pagination": pagination,
	})
}

func (u *userController) CreateUser(ctx *gin.Context) {
	response := utils.Response{C: ctx}
	Uvalidation := validation.NewCreateUserValidation()

	if err := Uvalidation.Bind(ctx); err != nil {
		ResponError := validator2.BindErrors(err)
		response.ResponseFormatter(http.StatusBadRequest, "Invalid Form", ResponError, nil)
		return
	}

	newUsr := model.NewUserModel()
	newUsr.Name = Uvalidation.Name
	newUsr.Email = Uvalidation.Email
	newUsr.Phone = Uvalidation.Phone
	newUsr.Address = Uvalidation.Address
	res, err := u.services.CreateUser(newUsr)

	if err != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "Failed Save Data", err.Error(), gin.H{
			"errors": err.Error(),
		})
	}

	response.ResponseFormatter(http.StatusOK, "Success Save Data", nil, gin.H{
		"email":   res.Email,
		"name":    res.Name,
		"phone":   res.Phone,
		"address": res.Address,
	})
}

func (u *userController) FindById(ctx *gin.Context) {
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
	response.ResponseFormatter(http.StatusOK, "User By Id", nil, responses)
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
	_ = u.services.Update(id, newUsr)

	updData := map[string]interface{}{}
	v := reflect.ValueOf(newUsr)
	typeOfV := v.Type()
	for i := 0; i < v.NumField(); i++ {
		updData[typeOfV.Field(i).Name] = v.Field(i).Interface()
	}
	delete(updData, "BaseModel")
	fmt.Println("updData", updData)

	response.ResponseFormatter(http.StatusOK, fmt.Sprintf("sukses update data with id = %s", strconv.Itoa(int(id))), nil, updData)
}
