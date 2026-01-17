package handlers

import (
	"github.com/gin-gonic/gin"
)

func ContextData(c *gin.Context) {
	requestID := c.GetString("RequestID")
	c.JSON(200, gin.H{
		"requestID": requestID,
	})
}
