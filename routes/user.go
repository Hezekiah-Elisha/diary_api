package routes

import (
	"diary_api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("/", controllers.GetUsers)
		users.POST("/", controllers.CreateUser)
		users.GET("/:id", controllers.GetUser)
		users.DELETE("/:id", controllers.DeleteUser)
		users.PUT("/:id", controllers.UpdateUser)
	}
}
