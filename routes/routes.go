package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	MainRoutes(r)
	UserRoutes(r)
}
