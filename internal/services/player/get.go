package player

import (
	"fmt"
	"log"

	"github.com/jairogloz/go-l/internal/domain"
)

type GetPlayerError struct {
	details error
}

func (e GetPlayerError) Error() string {
	return fmt.Sprintf("error getting player: %s", e.details.Error())
}

func (s *Service) Get(id string) (player domain.Player, err error) {
	player, err = s.Repo.Get(id)
	if err != nil {
		log.Println(err.Error())
		err = GetPlayerError{err}
		return
	}

	return
}
