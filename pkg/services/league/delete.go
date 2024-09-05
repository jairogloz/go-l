package league

import (
    "errors"
	"github.com/jairogloz/go-l/pkg/domain"
)

// Delete league by id
func (s *Service) Delete(id string) (err error) {
    if id == "" {
        return errors.New("id is required")
    }
    err = s.Repo.Delete(id)
    if err != nil {
        return domain.ManageError(err, "Error deleting league")
    }
    return nil
}
