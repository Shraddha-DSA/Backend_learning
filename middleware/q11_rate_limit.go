package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type client struct {
	count     int
	lastReset time.Time
}

var (
	mu      sync.Mutex
	clients = make(map[string]*client)
)

func RateLimit(maxRequests int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		mu.Lock()
		defer mu.Unlock()
		cl, exists := clients[ip]
		if !exists {
			clients[ip] = &client{
				count:     1,
				lastReset: time.Now(),
			}
			c.Next()
			return
		}
		if cl.count >= maxRequests {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
			})
			return
		}
		cl.count++
		c.Next()
	}
}
