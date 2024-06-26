package core

// LeagueCreateParams is a struct that represents the parameters needed to create a league.
type LeagueCreateParams struct {
	Description string `json:"description" binding:"required"`
	Name        string `json:"name" binding:"required"`
}
