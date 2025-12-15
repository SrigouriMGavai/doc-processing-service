package repository

import (
	"context"
	"time"

	"gitlab.com/srigourimgavai-group/doc-processing-service/internal/model"
	"github.com/go-kivik/kivik/v3"
)

type DocumentRepository struct {
	db *kivik.DB
}

func NewDocumentRepository(db *kivik.DB) *DocumentRepository {
	return &DocumentRepository{db: db}
}

func (r *DocumentRepository) Create(ctx context.Context, title string) (*model.Document, error) {
	doc := &model.Document{
		Title:     title,
		Status:    "created",
		CreatedAt: time.Now(),
	}

	id, _, err := r.db.CreateDoc(ctx, doc)
	if err != nil {
		return nil, err
	}

	doc.ID = id
	return doc, nil
}
