package team

import (
	"context"
	"log"
	"time"

	"github.com/jairogloz/go-l/internal/domain"
)

/*
.
*/
func (s *Service) Create(ctx context.Context, team *domain.Team) (err error) {
	now := time.Now().UTC()
	team.CreatedAt = &now

	err = s.Repo.Insert(ctx, team)
	if err != nil {
		appErr := domain.ManageError(err, "")
		log.Println(appErr.Error())
		return appErr
	}

	return nil
}
