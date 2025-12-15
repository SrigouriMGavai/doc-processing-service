package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/srigourimgavai-group/doc-processing-service/internal/repository"
)

type DocumentHandler struct {
	repo *repository.DocumentRepository
}

func NewDocumentHandler(repo *repository.DocumentRepository) *DocumentHandler {
	return &DocumentHandler{repo: repo}
}

func (h *DocumentHandler) CreateDocument(c *gin.Context) {
	var req struct {
		Title string `json:"title"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	doc, err := h.repo.Create(c.Request.Context(), req.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create document",
		})
		return
	}

	c.JSON(http.StatusCreated, doc)
}
func (h *DocumentHandler) GetDocumentByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "missing document id",
		})
		return
	}

	doc, err := h.repo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "document not found",
		})
		return
	}

	c.JSON(http.StatusOK, doc)
}
