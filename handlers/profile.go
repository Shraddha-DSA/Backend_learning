package handlers

import "github.com/gin-gonic/gin"

func Profile(c *gin.Context) {
	userID := c.GetInt("userID")
	c.JSON(200, gin.H{
		"message": "Authenticated",
		"user_id": userID,
	})
}
