package team

import (
	"context"

	"github.com/jairogloz/go-l/pkg/domain"
)

// Delete removes a team from the repository.
// It returns an error if the team does not exist or if there is an unexpected error.
func (s Service) Delete(ctx context.Context, id string) (err error) {

	err = s.Repo.Delete(ctx, id)

	if err != nil {
		return domain.ManageError(err, "unexpected error deleting team")
	}

	return nil
}
