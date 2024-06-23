package core

import "github.com/jairogloz/go-l/pkg/domain"

// PlayerCreateParams is a struct that represents the parameters needed to create a player.
type PlayerCreateParams struct {
	ContactInfo *domain.ContactInfo `json:"contact_info" `
	DateOfBirth *domain.Date        `json:"date_of_birth" binding:"required"`
	FirstName   string              `json:"first_name" binding:"required"`
	LastName    string              `json:"last_name" binding:"required"`
	TeamInfo    *domain.TeamInfo    `json:"team_info" binding:"required"`
}
