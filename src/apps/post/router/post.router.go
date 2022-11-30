package router

import (
	"github.com/gin-gonic/gin"
	"gorestapi/config"
	"gorestapi/src/apps/post/controller"
	"gorestapi/src/apps/post/repository"
	"gorestapi/src/apps/post/service"
	"net/http"
)

var (
	db             = config.SetUp()
	PostRepository = repository.NewPostRepository(db)
	PostService    = service.NewPostService(PostRepository)
	PostController = controller.NewPostController(PostService)
)

func PostRouter(router *gin.RouterGroup) {
	router.StaticFS("/public", http.Dir("public/"))
	router.GET("/all", PostController.GetAllPost)
	router.GET("/:id", PostController.FindById)
	router.POST("/create", PostController.CreatePost)
	router.POST("/:id", PostController.Update)
}
