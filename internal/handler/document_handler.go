package handler

import (
	"net/http"

	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gitlab.com/srigourimgavai-group/doc-processing-service/internal/model"
	"gitlab.com/srigourimgavai-group/doc-processing-service/internal/repository"

	
)

type DocumentHandler struct {
	repo  *repository.DocumentRepository
	cache *redis.Client
}

func NewDocumentHandler(repo *repository.DocumentRepository, cache *redis.Client) *DocumentHandler {
	return &DocumentHandler{
		repo:  repo,
		cache: cache,
	}
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

	ctx := c.Request.Context()
	cacheKey := "document:" + id

	// 1️⃣ Try Redis cache
	if cached, err := h.cache.Get(ctx, cacheKey).Result(); err == nil {
		var doc model.Document
		if err := json.Unmarshal([]byte(cached), &doc); err == nil {
			c.JSON(http.StatusOK, doc)
			return
		}
	}

	// 2️⃣ Fetch from CouchDB
	doc, err := h.repo.GetByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "document not found",
		})
		return
	}

	// 3️⃣ Store in Redis (TTL 1 min)
	if data, err := json.Marshal(doc); err == nil {
		h.cache.Set(ctx, cacheKey, data, time.Minute)
	}

	c.JSON(http.StatusOK, doc)
}
