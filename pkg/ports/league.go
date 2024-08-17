package ports

import (
	"context"

	"github.com/jairogloz/go-l/pkg/domain"
)

// LeagueService is the interface that have methods to interact with the league entity.
type LeagueService interface {
	Create(ctx context.Context, league *domain.League) (err error)
	Get(id string) (league *domain.League, err error)
}

// LeagueRepository is the interface that have methods to interact with the league entity in the database.
type LeagueRepository interface {
	Insert(ctx context.Context, league *domain.League) (err error)
	Get(id string) (league *domain.League, err error)
}
