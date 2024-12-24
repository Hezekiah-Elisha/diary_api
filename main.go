package main

import (
	"diary_api/config"
	"diary_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB()

	// defer db.Close()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
