package handlers

import (
	"go_assignment/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserJSON(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    req,
	})
}
