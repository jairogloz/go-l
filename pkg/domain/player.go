package domain

import "time"

// Player represents a player in the domain.
type Player struct {
	ContactInfo *ContactInfo `json:"contact_info" bson:"contact_info"`
	CreatedAt   *time.Time   `json:"created_at" bson:"created_at"`
	DateOfBirth *Date        `json:"date_of_birth"`
	FirstName   string       `json:"first_name" bson:"first_name"`
	ID          interface{}  `json:"id" bson:"_id"`
	LastName    string       `json:"last_name" bson:"last_name"`
	TeamInfo    *TeamInfo    `json:"team_info" bson:"team_info"`
	UpdatedAt   *time.Time   `json:"updated_at" bson:"updated_at"`
}
