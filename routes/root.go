package routes

import (
	"diary_api/controllers"

	"github.com/gin-gonic/gin"
)

func MainRoutes(r *gin.Engine) {
	r.GET("/", controllers.GetRoot)
}
