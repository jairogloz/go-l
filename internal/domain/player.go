package domain

import "time"

// Player represents a player in the domain.
type Player struct {
	ID          string       `json:"id" mongo:"_id"`
	FirstName   string       `json:"first_name" mongo:"first_name"`
	LastName    string       `json:"last_name" mongo:"last_name"`
	DateOfBirth *Date        `json:"age" binding:"required"`
	CreatedAt   *time.Time   `json:"created_at" mongo:"created_at"`
	UpdatedAt   *time.Time   `json:"updated_at" mongo:"updated_at"`
	ContactInfo *ContactInfo `json:"contact_info" mongo:"contact_info"`
	TeamInfo    *TeamInfo    `json:"team_info" mongo:"team_info"`
}
