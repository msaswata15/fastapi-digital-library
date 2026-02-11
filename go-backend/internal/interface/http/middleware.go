package http

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("startTime", start)
		ua := c.GetHeader("User-Agent")
		if ua == "" {
			ua = "Unknown"
		}
		log.Printf("[LOG] Request received from: %s", ua)
		c.Next()
		elapsed := time.Since(start)
		log.Printf("[LOG] Request processed in: %s", elapsed)
	}
}
