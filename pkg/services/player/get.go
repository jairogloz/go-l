package player

import (
	"errors"
	"fmt"
	"log"

	"github.com/jairogloz/go-l/pkg/domain"
)

// Get retrieves a player by their ID from the repository.
// It returns a domain.Player object and an error.
// If the provided ID is empty, it returns an error.
// If the player is not found in the repository, it returns a domain-specific not found error.
// If there is a timeout error when accessing the repository, it returns a domain-specific timeout error.
// For any other errors, it logs the error and returns a generic error.
func (s *Service) Get(id string) (player *domain.Player, err error) {

	if id == "" {
		return nil, errors.New("id is required")
	}

	player, err = s.Repo.Get(id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, domain.NewAppError(
				domain.ErrCodeNotFound,
				fmt.Sprintf("player with id '%s' not found", id))
		}
		if errors.Is(err, domain.ErrTimeout) {
			return nil, domain.NewAppError(
				domain.ErrCodeTimeout,
				"timeout error, try again later")
		}
		log.Println(err.Error())
		return nil, fmt.Errorf("unexpected error getting player: %w", err)
	}

	return player, nil
}
