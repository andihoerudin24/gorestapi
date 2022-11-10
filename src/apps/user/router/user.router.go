package router

import (
	"github.com/gin-gonic/gin"
	"gorestapi/config"
	"gorestapi/src/apps/user/controller"
	"gorestapi/src/apps/user/repository"
	"gorestapi/src/apps/user/service"
	"net/http"
)

var (
	db             = config.SetUp()
	UserRepository = repository.NewUserRepository(db)
	UserServices   = service.NewUserService(UserRepository)
	UserController = controller.NewUserController(UserServices)
)

func UserRouter(router *gin.RouterGroup) {
	r := *gin.Default()

	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.StaticFS("/public", http.Dir("public/"))
	router.GET("/all", UserController.GetAllUser)
	router.POST("/create", UserController.CreateUser)
	router.GET("/:id", UserController.FindById)
	router.POST("/:id", UserController.Update)
}
