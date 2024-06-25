package domain

import "time"

// Tournament reflects a sports tournament in a League in our system.
// A tournament can have one or more Editions, for instance, the League Golang League
// can have a Tournament called Kids Tournament, and this Tournament can have an Edition
// called 2024 Edition, then next year the same Tournament can have a new Edition called
// 2025 Edition.
type Tournament struct {
	ID          interface{} `json:"id" bson:"_id,omitempty"`
	Name        string      `json:"name" bson:"name"`
	Description string      `json:"description" bson:"description"`
	URL         string      `json:"url" bson:"url"`
	CreatedAt   *time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt   *time.Time  `json:"updated_at" bson:"updated_at"`
}
