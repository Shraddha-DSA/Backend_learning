package handlers

import (
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "GET",
		"path":   "/users",
	})
}
func CreateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "POST",
		"path":   "/users",
	})
}
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"method": "PUT",
		"path":   "/users/" + id,
	})
}
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"method": "DELETE",
		"path":   "/users/" + id,
	})
}
