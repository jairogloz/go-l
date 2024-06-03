package team

import (
	"github.com/jairogloz/go-l/internal/ports"
)

// Make sure Repository implements ports.TeamRepository
// at compile time
var _ ports.TeamRepository = &Repository{}

type Repository struct {
}
