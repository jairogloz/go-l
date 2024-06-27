package league

import (
	"github.com/jairogloz/go-l/pkg/ports"
)

// Make sure Service implements the LeagueService interface
// at compile time.
var _ ports.LeagueService = &Service{}

type Service struct {
	Repo ports.LeagueRepository
}
