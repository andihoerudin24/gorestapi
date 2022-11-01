package controller

import (
	"github.com/gin-gonic/gin"
	"gorestapi/src/apps/user/model"
	"gorestapi/src/apps/user/service"
	"gorestapi/src/apps/user/validation"
	"gorestapi/utils"
	validator2 "gorestapi/validator"
	"net/http"
)

type UserController interface {
	GetAllUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
}

type userController struct {
	services service.UserService
}

func NewUserController(services service.UserService) *userController {
	return &userController{services: services}
}

func (u *userController) GetAllUser(ctx *gin.Context) {
	response := utils.Response{C: ctx}
	DataUser := u.services.GetAllUser()
	if DataUser == nil {
		response.ResponseFormatter(http.StatusNotFound, "data not found", DataUser, nil)
	}
	response.ResponseFormatter(http.StatusOK, "List User", nil, gin.H{
		"data": DataUser,
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
