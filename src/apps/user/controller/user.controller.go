package controller

import (
	"github.com/gin-gonic/gin"
	"gorestapi/src/apps/user/service"
	"gorestapi/utils"
	"net/http"
)

type UserController interface {
	GetAllUser(ctx *gin.Context)
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
		ctx.JSON(http.StatusFound, gin.H{
			"code":    http.StatusFound,
			"message": "data not found",
		})
	}
	response.ResponseFormatter(http.StatusOK, "List User", nil, gin.H{
		"data": DataUser,
	})
}
