package player

import (
	"github.com/jairogloz/go-l/pkg/domain"
)

// Delete player by id
func (s *Service) Delete(id string) (err error) {
	err = s.Repo.Delete(id)
	if err != nil {
		return domain.ManageError(err, "Error deleting player")
	}
	return nil
}
