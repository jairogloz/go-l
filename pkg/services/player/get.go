package player

import (
	"errors"
	"fmt"
	"log"

	"github.com/jairogloz/go-l/pkg/domain"
)

func (s *Service) Get(id string) (player *domain.Player, err error) {
	player, err = s.Repo.Get(id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, domain.NewAppError(
				domain.ErrCodeNotFound,
				fmt.Sprintf("player with id '%s' not found", id))
		}
		log.Println(err.Error())
		return nil, fmt.Errorf("unexpected error getting player: %w", err)
	}

	return player, nil
}
