package controlleruser

import (
	"github.com/gin-gonic/gin"
	user2 "gorestapi/services/user"
	"net/http"
)

type ControllerUser interface {
	GetAllUser(ctx *gin.Context)
}

type controllerUser struct {
	services user2.ServicesUser
}

func NewControllerUser(services user2.ServicesUser) *controllerUser {
	return &controllerUser{services: services}
}

func (c *controllerUser) GetAllUser(ctx *gin.Context) {
	datauser := c.services.GetAllUser()
	if datauser == nil {
		ctx.JSON(http.StatusFound, gin.H{
			"code":    http.StatusFound,
			"message": "data not found",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": datauser,
	})
}
