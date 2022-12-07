package router

import (
	"github.com/gin-gonic/gin"
	"gorestapi/cache"
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
	InitRedis      = config.InitRedis()
	RedisCache     = cache.NewRedisCache(InitRedis)
	UserController = controller.NewUserController(UserServices, RedisCache)
)

func UserRouter(router *gin.RouterGroup) {
	r := *gin.Default()

	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.StaticFS("/public", http.Dir("public/"))
	router.GET("/all", UserController.GetAllUser)
	router.POST("/create", UserController.CreateUser)
	router.GET("/:id", UserController.FindById)
	router.POST("/:id", UserController.Update)
	router.DELETE("/:id", UserController.Delete)
}
