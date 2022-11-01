package routes

import (
	"github.com/gin-gonic/gin"
	"gorestapi/src/apps/user/router"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/api/v1")
	{
		router.UserRouter(v1.Group("/user"))
	}

	return r
}
