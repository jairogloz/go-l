package tournament

import "github.com/jairogloz/go-l/pkg/ports"

// Make sure Service implements the TournamentService interface
// at compile time.
var _ ports.TournamentService = &Service{}

// Service is a struct that represents the service for the tournament entity.
type Service struct {
	Repo  ports.TournamentRepository
	Clock ports.Clock
}
