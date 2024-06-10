package team

import (
	"context"
	"fmt"
	"log"

	"github.com/jairogloz/go-l/internal/domain"
)

type GetTeamError struct {
	details error
}

func (e GetTeamError) Error() string {
	return fmt.Sprintf("error getting team: %s", e.details.Error())
}

func (s Service) Get(ctx context.Context, id string) (team domain.Team, players []domain.Player, err error) {

	team, err = s.Repo.Get(ctx, id)
	if err != nil {
		log.Println(err.Error())
		err = GetTeamError{err}
		return
	}

	players, err = s.Repo.GetPlayers(ctx, id)
	if err != nil {
		log.Println(err.Error())
		err = GetTeamError{err}
		return
	}

	return
}
