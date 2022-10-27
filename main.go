package main

import (
	"github.com/gin-gonic/gin"
	"gorestapi/config"
	"gorestapi/routes"
	"os"
)

var (
	db = config.SetUp()
)

func main() {
	defer config.CloseDatabase(db)

	routeMain := gin.Default()
	routes.RouteInit(routeMain.Group("rest"))
	routeMain.Run(os.Getenv("APP_URL") + ":" + os.Getenv("APP_PORT"))
}
