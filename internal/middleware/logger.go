package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		latency := time.Since(start)
		method := c.Request.Method
		path := c.Request.URL.Path
		status := c.Writer.Status()

		// Simple log (stdout for now)
		println(
			"method:", method,
			"path:", path,
			"status:", status,
			"latency:", latency.String(),
		)
	}
}
