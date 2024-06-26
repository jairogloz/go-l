package team

import (
	"github.com/jairogloz/go-l/pkg/ports"
)

// Make sure Service implements the TeamService interface
// at compile time.
var _ ports.TeamService = &Service{}

type Service struct {
	Repo ports.TeamRepository
}
