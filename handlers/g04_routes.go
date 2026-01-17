package handlers

import (
	"github.com/gin-gonic/gin"
)

func RouteParams(c *gin.Context) {
	userID := c.Param("id")
	orderID := c.Param("orderId")
	c.JSON(200, gin.H{
		"user_id":  userID,
		"order_id": orderID,
	})

}
