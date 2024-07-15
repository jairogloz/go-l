package team

import (
	"context"

	"github.com/jairogloz/go-l/pkg/domain"
)

func (s Service) Delete(ctx context.Context, id string) (err error) {

	err = s.Repo.Delete(ctx, id)

	if err != nil {
		return domain.ManageError(err, "unexpected error deleting team")
	}

	return nil
}
