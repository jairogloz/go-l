package player

import (
	"github.com/jairogloz/go-l/internal/ports"
)

// Make sure Service implements the PlayerService interface
// at compile time.
var _ ports.PlayerService = &Service{}

type Service struct {
	Repo ports.PlayerRepository
}
