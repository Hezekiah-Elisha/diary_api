package routes

import (
	"diary_api/controllers"
	"diary_api/middleware"

	"github.com/gin-gonic/gin"
)

func UserSetupRoutes(r *gin.Engine) {
	// Create a route group for user routes
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.GET("/:id", controllers.GetUser)
	}

	// Create a route group for user routes that require authentication
	authUserRoutes := r.Group("/users")
	authUserRoutes.Use(middleware.AuthMiddleware())
	{
		authUserRoutes.POST("/", controllers.CreateUser)
		authUserRoutes.PUT("/:id", controllers.UpdateUser)
		authUserRoutes.DELETE("/", controllers.DeleteAllUsers)
		authUserRoutes.DELETE("/:id", controllers.DeleteUser)
	}
}
