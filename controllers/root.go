package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Hello, welcome to the diary API",
	})
}
