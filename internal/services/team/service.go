package team

import (
	"github.com/jairogloz/go-l/internal/ports"
)

// Make sure Service implements the TeamService interface
// at compile time.
var _ ports.TeamService = &Service{}

type Service struct {
	Repo ports.TeamRepository
}
