package router

import (
	"github.com/gin-gonic/gin"
	"gorestapi/config"
	"gorestapi/src/apps/post/controller"
	"gorestapi/src/apps/post/repository"
	"gorestapi/src/apps/post/service"
)

var (
	db             = config.SetUp()
	PostRepository = repository.NewPostRepository(db)
	PostService    = service.NewPostService(PostRepository)
	PostController = controller.NewPostController(PostService)
)

func PostRouter(router *gin.RouterGroup) {
	router.GET("/all", PostController.GetAllPost)
	router.POST("/create", PostController.CreatePost)
}
