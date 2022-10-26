package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorestapi/routes"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("failed load file env")
	}
	routeMain := gin.Default()
	routes.RouteInit(routeMain.Group("rest"))
	routeMain.Run(os.Getenv("APP_URL") + ":" + os.Getenv("APP_PORT"))
}
