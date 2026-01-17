package handlers

import (
	"go_assignment/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequests struct {
	UserID int `json:"user_id"`
}

func Logins(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Login request",
		})
		return
	}
	token, err := utils.GenerateToken(req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "token generation failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
