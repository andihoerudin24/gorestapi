package router

import (
	"github.com/gin-gonic/gin"
	redis2 "github.com/go-redis/redis/v9"
	"gorestapi/cache"
	"gorestapi/config"
	"gorestapi/src/apps/post/controller"
	"gorestapi/src/apps/post/repository"
	"gorestapi/src/apps/post/service"
	"net/http"
)

var (
	redis          *redis2.Client
	db             = config.SetUp()
	PostRepository = repository.NewPostRepository(db)
	PostService    = service.NewPostService(PostRepository)
	InitRedis      = config.InitRedis()
	redisCache     = cache.NewRedisCache(InitRedis)
	PostController = controller.NewPostController(PostService, redisCache)
)

func PostRouter(router *gin.RouterGroup) {
	router.StaticFS("/public", http.Dir("public/"))
	router.GET("/all", PostController.GetAllPost)
	router.GET("/:id", PostController.FindById)
	router.POST("/create", PostController.CreatePost)
	router.POST("/:id", PostController.Update)
	router.DELETE("/:id", PostController.Delete)
}
