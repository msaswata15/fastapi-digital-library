package http

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		ua := c.GetHeader("User-Agent")
		if ua == "" {
			ua = "Unknown"
		}
		log.Printf("[LOG] Request received from: %s", ua)
		c.Next()
		elapsed := time.Since(start)
		c.Writer.Header().Set("X-Process-Time", elapsed.String())
	}
}
