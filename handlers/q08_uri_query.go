package handlers

import (
	"go_assignment/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserProfile(c *gin.Context) {
	var uri models.UserURI
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid user id",
		})
		return
	}
	var query models.ProfileQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid query parameters",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user_id": uri.ID,
		"verbose": query.Verbose,
		"lang":    query.Lang,
	})
}
