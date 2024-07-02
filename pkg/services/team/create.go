package team

import (
	"context"
	"time"

	"github.com/jairogloz/go-l/pkg/domain"
)

// Create creates a new team based on the provided data and stores it in the database.
// Method: Create
// Path: pkg/services/team/create.go
// Input: ctx context.Context, team *domain.Team
// Output: error

func (s Service) Create(ctx context.Context, team *domain.Team) (err error) {
	now := time.Now().UTC()
	team.CreatedAt = &now

	err = s.Repo.Insert(ctx, team)
	if err != nil {
		return domain.ManageError(err, "Error creating team")
	}
	return nil
}
