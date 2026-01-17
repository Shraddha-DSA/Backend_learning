package handlers

import (
	"github.com/gin-gonic/gin"
)

func ProtectedRoute(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "You are authorized",
	})
}
