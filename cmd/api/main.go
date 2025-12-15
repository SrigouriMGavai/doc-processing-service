package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/srigourimgavai-group/doc-processing-service/internal/config"
	"gitlab.com/srigourimgavai-group/doc-processing-service/internal/db"
	"gitlab.com/srigourimgavai-group/doc-processing-service/internal/handler"
	"gitlab.com/srigourimgavai-group/doc-processing-service/internal/middleware"
	"gitlab.com/srigourimgavai-group/doc-processing-service/internal/repository"
)

func main() {
	// Create a Gin router
	cfg := config.Load()
	couch, err := db.ConnectCouchDB()
	//_, err := db.ConnectCouchDB()

	if err != nil {
		panic(err)
	}
	println("Connected to CouchDB")
	docRepo := repository.NewDocumentRepository(couch.DB)
	docHandler := handler.NewDocumentHandler(docRepo)

	//println("Test document created with ID:", doc.ID)

	r := gin.Default()
	r.Use(middleware.RequestLogger())

	// Simple health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	r.POST("/documents", docHandler.CreateDocument)
	r.GET("/documents/:id", docHandler.GetDocumentByID)

	// Start HTTP server on port 8080
	r.Run(":" + cfg.AppPort)

}
