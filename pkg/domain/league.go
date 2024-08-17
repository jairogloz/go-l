package domain

import "time"

// League reflects a sports league in our system.
// A league can have one or more Tournaments, and a Tournament can have one or more
// Editions, for instance, the League Golang League and have a Tournament called
// Kids Tournament, and this Tournament can have an Edition called 2024 Edition,
// then next year the same Tournament can have a new Edition called 2025 Edition.
type League struct {
	CreatedAt   *time.Time  `json:"created_at" bson:"created_at"`
	Description string      `json:"description" bson:"description"`
	ID          interface{} `json:"id" bson:"_id,omitempty"`
	Name        string      `json:"name" bson:"name"`
	UpdatedAt   *time.Time  `json:"updated_at" bson:"updated_at"`
}
