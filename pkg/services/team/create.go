package team

import (
	"context"

	"github.com/jairogloz/go-l/pkg/domain"
)

func (s Service) Create(ctx context.Context, team *domain.Team) (err error) {
	err = s.Repo.Insert(ctx, team)
	if err != nil {
		return domain.ManageError(err, "Error creating team")
	}
	return nil
}
