package player

import (
	"log"

	"github.com/jairogloz/go-l/internal/domain"
)

// Delete player by id
func (s *Service) Delete(id string) (err error) {
	err = s.Repo.Delete(id)
	if err != nil {
		return manageError(err)
	}
	return nil
}

func manageError(err error) error {
	switch err {
	case domain.ErrIncorrectID:
		log.Println("incorrect id error")
		appErr := domain.AppError{
			Code: domain.ErrCodeInvalidParams,
			Msg:  "error deleting player: incorrect id",
		}
		return appErr
	case domain.ErrNotFound:
		log.Println("player not found error")
		appErr := domain.AppError{
			Code: domain.ErrCodeNotFound,
			Msg:  "error deleting player: player not found",
		}
		return appErr
	default:
		log.Println(err.Error())
		appErr := domain.AppError{
			Code: domain.ErrCodeInternalServerError,
			Msg:  "error deleting player",
		}
		return appErr
	}
}
