package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/srigourimgavai-group/doc-processing-service/internal/config"
	"gitlab.com/srigourimgavai-group/doc-processing-service/internal/middleware"
)

func main() {
	// Create a Gin router
	cfg := config.Load()

	r := gin.Default()
	r.Use(middleware.RequestLogger())

	// Simple health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Start HTTP server on port 8080
	r.Run(":" + cfg.AppPort)

}
