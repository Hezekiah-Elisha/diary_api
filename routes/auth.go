package routes

import (
	"diary_api/controllers"

	"github.com/gin-gonic/gin"
)

func AuthSetupRoutes(r *gin.Engine) {
	r.POST("/auth/login", controllers.HandleLogin)
	r.POST("/auth/logout", controllers.HandleLogout)
	r.POST("/auth/register", controllers.HandleRegister)
}
