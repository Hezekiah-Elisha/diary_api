package routes

import (
	"diary_api/controllers"
	"diary_api/middleware"

	"github.com/gin-gonic/gin"
)

func EntrySetupRoutes(r *gin.Engine) {
	entryRoutes := r.Group("/entries")
	{
		entryRoutes.GET("/", controllers.GetEntries)
		entryRoutes.GET("/:id", controllers.GetEntry)
		entryRoutes.POST("/", controllers.CreateEntry)
		// entryRoutes.PUT("/:id", controllers.UpdateEntry)
		entryRoutes.DELETE("/:id", controllers.DeleteEntry)
	}
	entryRoutes.Use(middleware.AuthMiddleware())
}
