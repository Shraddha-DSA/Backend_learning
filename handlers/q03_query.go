package handlers

import (
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	query := c.Query("query")
	page := c.DefaultQuery("page", "1")
	c.JSON(200, gin.H{
		"query": query,
		"page":  page,
	})
}
