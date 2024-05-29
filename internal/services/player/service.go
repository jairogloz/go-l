package player

import (
	"github.com/jairogloz/go-l/internal/ports"
)

type Service struct {
	Repo ports.PlayerRepository
}
