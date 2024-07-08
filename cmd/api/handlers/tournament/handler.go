package tournament

import "github.com/jairogloz/go-l/pkg/ports"

// Handler is a struct that represents the handler for the tournament entity.
type Handler struct {
	TournamentService ports.TournamentService
}
