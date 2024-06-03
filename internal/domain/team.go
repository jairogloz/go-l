package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// Team reflects a sports team in our system.
type Team struct {
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
	ID        interface{}        `json:"-" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at"`
}

// TeamInfo represents the team information of a player, not the
// full team information.
type TeamInfo struct {
	TeamID       string `json:"team_id" bson:"team_id"`
	JerseyNumber int    `json:"jersey_number" bson:"jersey_number"`
}
