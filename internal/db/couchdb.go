package db

import (
	"context"
	"fmt"

	"github.com/go-kivik/couchdb/v3"
	"github.com/go-kivik/kivik/v3"
)

type CouchDB struct {
	Client *kivik.Client
	DB     *kivik.DB
}

func ConnectCouchDB() (*CouchDB, error) {
	ctx := context.Background()

	// Create CouchDB client
	client, err := kivik.New("couch", "http://127.0.0.1:5984/")
	if err != nil {
		return nil, fmt.Errorf("failed to create couchdb client: %w", err)
	}

	// ✅ CORRECT authentication for kivik v3
	if err := client.Authenticate(ctx, couchdb.BasicAuth("admin", "admin")); err != nil {
		return nil, fmt.Errorf("couchdb authentication failed: %w", err)
	}

	// Connect to database
	db := client.DB(ctx, "documents")
	if err := db.Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &CouchDB{
		Client: client,
		DB:     db,
	}, nil
}
