package middleware

import (
	"go_assignment/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Errorhandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			apiErr, ok := err.Meta.(models.APIError)
			if !ok {
				apiErr = models.APIError{
					Code:    http.StatusInternalServerError,
					Message: "Internal server error",
				}
			}
			c.JSON(apiErr.Code, apiErr)

		}
	}

}
