package core

// TournamentCreateParams is a struct that represents the parameters needed to create a tournament.
type TournamentCreateParams struct {
	Description string `json:"description" binding:"required"`
	Name        string `json:"name" binding:"required"`
	URL         string `json:"url" binding:"required"`
}
