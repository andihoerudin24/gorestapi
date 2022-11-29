package routes

import (
	"github.com/gin-gonic/gin"
	postrouter "gorestapi/src/apps/post/router"
	userrouter "gorestapi/src/apps/user/router"
)

func Router() *gin.Engine {
	r := gin.Default()
	//r.ForwardedByClientIP = true
	//r.Use(utils.Logger_JSON("logger.log", true))

	v1 := r.Group("/api/v1")
	{
		userrouter.UserRouter(v1.Group("/user"))
		postrouter.PostRouter(v1.Group("/post"))

	}
	return r
}
