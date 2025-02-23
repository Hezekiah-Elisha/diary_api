package routes

import (
	"diary_api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// r.Group("/api")

	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.PUT("/users/:id", controllers.UpdateUser)
}
