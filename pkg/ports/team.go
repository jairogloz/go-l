package ports

import (
	"context"

	"github.com/jairogloz/go-l/pkg/domain"
)

type TeamService interface {
	Create(ctx context.Context, team *domain.Team) (err error)
	Get(ctx context.Context, id string) (team *domain.Team, err error)
	Delete(ctx context.Context, id string) (err error)
}

type TeamRepository interface {
	Insert(ctx context.Context, team *domain.Team) (err error)
	Get(ctx context.Context, id string) (team *domain.Team, err error)
	Delete(ctx context.Context, id string) (err error)
}
