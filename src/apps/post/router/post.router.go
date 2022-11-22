package router

import (
	"github.com/gin-gonic/gin"
	"gorestapi/config"
	"gorestapi/src/apps/post/controller"
	"gorestapi/src/apps/post/repository"
)

var (
	db             = config.SetUp()
	PostRepository = repository.NewPostRepository(db)
	PostController = controller.NewPostController(PostRepository)
)

func PostRouter(router *gin.RouterGroup) {
	router.GET("/all", PostController.GetAllPost)
}
