package team

import (
	"context"

	"github.com/jairogloz/go-l/pkg/domain"
)

// Get retrieves a team and its players.
// returns the team and its players or an error if something goes wrong.
func (s Service) Get(ctx context.Context, id string) (team *domain.Team, players []domain.Player, err error) {
	team, err = s.Repo.Get(ctx, id)

	if err != nil {
		return nil, nil, domain.ManageError(err, "unexpected error getting team")
	}

	players, err = s.Repo.GetPlayers(ctx, team.ID.(string))
	if err != nil {
		return nil, nil, domain.ManageError(err, "unexpected error getting players")
	}

	return team, players, nil
}
