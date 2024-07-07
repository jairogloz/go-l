package tournament

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jairogloz/go-l/pkg/domain"
)

// Delete deletes a tournament by id.
func (s *Service) Delete(ctx context.Context, id string) (err error) {
	err = s.Repo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrIncorrectID) {
			log.Println("Incorrect ID error")
			appErr := domain.AppError{
				Code: domain.ErrCodeInvalidParams,
				Msg:  "error deleting tournament: incorrect ID",
			}
			return appErr
		}

		if errors.Is(err, domain.ErrNotFound) {
			log.Println("Not found error")
			appErr := domain.AppError{
				Code: domain.ErrCodeNotFound,
				Msg:  "error deleting tournament: not found",
			}
			return appErr
		}

		log.Println(err.Error())
		return fmt.Errorf("error creating tournament: %w", err)
	}

	return
}
