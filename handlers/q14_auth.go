package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"go_assignment/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	UserID int `json:"user_id"`
}

func generateToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request",
		})
		return
	}

	token := generateToken()
	store.SaveToken(token, req.UserID)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
