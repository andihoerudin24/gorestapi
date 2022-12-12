package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	postrouter "gorestapi/src/apps/post/router"
	userrouter "gorestapi/src/apps/user/router"
	"gorestapi/utils"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.ForwardedByClientIP = true
	r.Use(utils.Logger_JSON("logger.log", true))
	r.LoadHTMLFiles("./src/view/swagger/index.html")
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "DELETE"},
		AllowHeaders:    []string{"Origin"},
		ExposeHeaders:   []string{"Content-Length"},
	}))
	v1 := r.Group("/api/v1")
	{
		v1.StaticFS("/public/swagger", http.Dir("public/swagger"))
		v1.GET("/", func(context *gin.Context) {
			context.HTML(http.StatusOK, "index.html", nil)
		})
		userrouter.UserRouter(v1.Group("/user"))
		postrouter.PostRouter(v1.Group("/post"))

	}
	return r
}
