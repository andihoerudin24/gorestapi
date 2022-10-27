package routes

import (
	"github.com/gin-gonic/gin"
	"gorestapi/config"
	"gorestapi/controller/controlleruser"
	"gorestapi/repository/user"
	user2 "gorestapi/services/user"
)

var (
	db             = config.SetUp()
	userrepository = user.NewUserRepository(db)
	userservices   = user2.NewServiceUser(userrepository)
	usercontroller = controlleruser.NewControllerUser(userservices)
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1/user")
	{
		v1.GET("/all", usercontroller.GetAllUser)
	}

	return r
}
