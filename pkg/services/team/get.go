package team

import (
	"context"

	"github.com/jairogloz/go-l/pkg/domain"
)

// Get retrieves a team from the database, or returns an error if something goes wrong.
func (s Service) Get(ctx context.Context, id string) (team *domain.Team, err error) {
	team, err = s.Repo.Get(ctx, id)

	if err != nil {
		return nil, domain.ManageError(err, "unexpected error getting team")
	}

	return team, nil
}
