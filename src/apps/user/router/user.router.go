package router

import (
	"github.com/gin-gonic/gin"
	"gorestapi/config"
	"gorestapi/src/apps/user/controller"
	"gorestapi/src/apps/user/repository"
	"gorestapi/src/apps/user/service"
)

var (
	db             = config.SetUp()
	UserRepository = repository.NewUserRepository(db)
	UserServices   = service.NewUserService(UserRepository)
	UserController = controller.NewUserController(UserServices)
)

func UserRouter(router *gin.RouterGroup) {
	router.GET("/all", UserController.GetAllUser)
}
