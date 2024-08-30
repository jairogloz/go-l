package league

import (
	"github.com/jairogloz/go-l/pkg/domain"
)

// Delete league by id
func (s *Service) Delete(id string) (err error) {
    err = s.Repo.Delete(id)
    if err != nil {
        return domain.ManageError(err, "Error deleting league")
    }
    return nil
}
