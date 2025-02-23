package main

import (
	"diary_api/config"
	"diary_api/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	// Set Gin mode from environment variable
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode // Default to debug mode if not set
	}
	gin.SetMode(ginMode)

	r := gin.Default()

	config.ConnectDB()

	routes.SetupRoutes(r)

	r.Run(":8081")
}
