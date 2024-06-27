package ports

import (
	"github.com/jairogloz/go-l/pkg/domain"
)

type LeagueService interface {
	Get(id string) (league *domain.League, err error)
}

type LeagueRepository interface {
	Get(id string) (league *domain.League, err error)
}
