package player

import (
	"errors"
	"fmt"
	"github.com/jairogloz/go-l/pkg/domain"
	"log"
	"time"
)

// Create is a method of the Service struct that creates a new player in the database.
//
// Params:
//   - player: A pointer to a domain.Player struct. This allows the method to modify the original Player struct.
//
// The method sets the CreatedAt field of the Player struct to the current time and then calls the Insert method of the Repo field of the Service struct, passing the Player struct as an argument.
//
// If the Insert method returns an error, the Create method checks if the error is a duplicate key error. If it is, it logs the error and returns a new AppError with the duplicate key error code and message. If the error is not a duplicate key error, it logs the error and returns a new error wrapping the original error with a message indicating that there was an error creating the player.
//
// If the Insert method does not return an error, the Create method sets the ID field of the Player struct to the ID returned by the Insert method. This allows the caller of the Create method to access the ID of the newly created Player.
//
// Return values:
//   - err: An error that will be nil if the Player was successfully created. If there was an error, it will be an error object describing the failure.
func (s *Service) Create(player *domain.Player) (err error) {
	now := time.Now().UTC()
	player.CreatedAt = &now

	err = s.Repo.Insert(player)
	if err != nil {
		if errors.Is(err, domain.ErrDuplicateKey) {
			log.Println("Duplicate key error")
			appErr := domain.AppError{
				Code: domain.ErrCodeDuplicateKey,
				Msg:  "error creating player: duplicate key error",
			}
			return appErr
		}
		log.Println(err.Error())
		return fmt.Errorf("error creating player: %w", err)
	}

	return nil
}
