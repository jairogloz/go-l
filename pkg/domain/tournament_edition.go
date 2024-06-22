package domain

import "time"

const (
	TournamentTypeLeague  = "league"   // liga, ej. La Liga en España o la Premier League en Inglaterra
	TournamentTypeCup     = "knockout" // eliminatorias, ej. Copa del Rey en España o FA Cup en Inglaterra
	TournamentTypePlayoff = "playoff"  // playoff, Etapa regular seguido de playoff, ej. Mundial de la Fifa
)

// TournamentEdition reflects a sports tournament edition in our system.
// An edition is a specific instance of a tournament, for instance, the League
// Golang League can have a Tournament called Kids Tournament, and this Tournament
// can have a TournamentEdition called 2024 Edition, then next year the same Tournament can
// have a new TournamentEdition called 2025 Edition.
//
// A TournamentEdition can have multiple teams that participate in it, they might vary between
// editions.
type TournamentEdition struct {
	CreatedAt     *time.Time     `json:"created_at" bson:"created_at"`
	Description   string         `json:"description" bson:"description"`
	EndDate       *time.Time     `json:"end_date" bson:"end_date"`
	GameDays      []time.Weekday `json:"game_days" bson:"game_days"`
	GameTimeRange *TimeRange     `json:"game_time_range" bson:"game_time_range"`
	ID            string         `json:"id" bson:"_id,omitempty"`
	Name          string         `json:"name" bson:"name"`
	StartDate     *time.Time     `json:"start_date" bson:"start_date"`
	Type          string         `json:"type" bson:"type"`
	UpdatedAt     *time.Time     `json:"updated_at" bson:"updated_at"`
}
