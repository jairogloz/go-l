package domain

// TeamInfo represents the team information of a player, not the
// full team information.
type TeamInfo struct {
	TeamID       string `json:"team_id" bson:"team_id"`
	JerseyNumber int    `json:"jersey_number" bson:"jersey_number"`
}
