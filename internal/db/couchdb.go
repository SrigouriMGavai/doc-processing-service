package db

import (
	"context"
	"fmt"

	"github.com/go-kivik/kivik/v3"
	_ "github.com/go-kivik/couchdb/v3"
)

type CouchDB struct {
	Client *kivik.Client
	DB     *kivik.DB
}

func ConnectCouchDB() (*CouchDB, error) {
	// CouchDB connection URL
	couchURL := "http://admin:admin@127.0.0.1:5984/"

	// Create client
	client, err := kivik.New("couch", couchURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create couchdb client: %w", err)
	}

	// Connect to database
	db := client.DB(context.Background(), "documents")

	// Check if DB exists
	if err := db.Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &CouchDB{
		Client: client,
		DB:     db,
	}, nil
}
