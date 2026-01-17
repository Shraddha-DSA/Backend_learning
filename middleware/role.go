package middleware

import (
	"github.com/gin-gonic/gin"
)

func RoleBased(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetHeader("X-Role")
		if role == "" {
			c.AbortWithStatusJSON(403, gin.H{
				"error": "Role not provided",
			})
			return
		}
		for _, r := range allowedRoles {
			if role == r {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(403, gin.H{
			"error": "access denied",
		})
	}
}
