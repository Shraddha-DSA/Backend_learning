package handlers

import (
	"net/http"
	"strconv"

	"go_assignment/models"
	"go_assignment/store"

	"github.com/gin-gonic/gin"
)

// CREATE USERS
func CreateUsers(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.ID = store.NextID
	store.NextID++

	store.Users[user.ID] = user

	c.JSON(http.StatusCreated, user)
}

// GET ALL USERS
func GetUserss(c *gin.Context) {
	users := make([]models.User, 0)

	for _, u := range store.Users {
		users = append(users, u)
	}

	c.JSON(http.StatusOK, users)
}

// GET USER BY ID
func GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}

	user, exists := store.Users[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UPDATE USERS
func UpdateUsers(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, exists := store.Users[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	user.Name = input.Name
	user.Email = input.Email

	store.Users[id] = user

	c.JSON(http.StatusOK, user)
}

// DELETE USER
func DeleteUsers(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}

	if _, exists := store.Users[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	delete(store.Users, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
	})
}
