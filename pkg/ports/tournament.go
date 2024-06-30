package ports

import (
	"context"

	"github.com/jairogloz/go-l/pkg/domain"
)

// TournamentService is the interface that have methods to interact with the tournament entity.
type TournamentService interface {
	Create(ctx context.Context, tournament *domain.Tournament) (err error)
}

// TournamentRepository is the interface that have methods to interact with the tournament entity in the database.
type TournamentRepository interface {
	Insert(ctx context.Context, tournament *domain.Tournament) (err error)
}
