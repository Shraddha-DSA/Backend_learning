package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ContextExample(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	resultChain := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		resultChain <- "Operation Completed"
	}()
	select {
	case <-ctx.Done():
		c.JSON(http.StatusRequestTimeout, gin.H{
			"error": "Request timed out",
		})
		return
	case result := <-resultChain:
		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
		return
	}
}
