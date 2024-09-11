package player

import (
	"errors"
	"fmt"
	"github.com/jairogloz/go-l/pkg/domain"
	"log"
)

// GetAll retrieves all players from the repository.
// It returns an Array of domain.Player objects and an error.
// If the player is not found in the repository, it returns a domain-specific not found error.
// If there is a timeout error when accessing the repository, it returns a domain-specific timeout error.
// For any other errors, it logs the error and returns a generic error.
func (s *Service) GetAll() (player []*domain.Player, err error) {

	player, err = s.Repo.GetAll()
	if err != nil {
		if errors.Is(err, domain.ErrTimeout) {
			return nil, domain.NewAppError(
				domain.ErrCodeTimeout,
				"timeout error, try again later")
		}
		log.Println(err.Error())
		return nil, fmt.Errorf("unexpected error getting players: %w", err)
	}

	return player, nil
}
