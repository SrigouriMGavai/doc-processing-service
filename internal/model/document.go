package model

import "time"

type Document struct {
	ID        string    `json:"id,omitempty" couchdb:"_id,omitempty"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
