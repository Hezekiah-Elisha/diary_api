package main

import (
	"diary_api/config"
	"diary_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	config.ConnectDB()

	routes.SetupRoutes(r)

	r.Run(":8081")
}
