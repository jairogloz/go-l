package core

import (
	"github.com/jairogloz/go-l/pkg/domain"
	"time"
)

// TournamentEditionParams is a struct that represents the parameters needed to edit a tournament.
type TournamentEditionParams struct {
	Description   string            `json:"description" binding:"required"`
	EndDate       *time.Time        `json:"end_date"` // Should be a valid date in the future, after than StartDate.
	GameDays      []time.Weekday    `json:"game_days" binding:"required"`
	GameTimeRange *domain.TimeRange `json:"game_time_range" bson:"game_time_range"`
	Name          string            `json:"name" binding:"required"`
	StartDate     *time.Time        `json:"start_date" bson:"start_date"`
	Type          string            `json:"type" bson:"type"`
	URL           string            `json:"url" binding:"required"`
}
