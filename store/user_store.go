package store

import (
	"go_assignment/models"
)

var (
	Users  = make(map[int]models.User)
	NextID = 1
)
