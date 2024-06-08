package team

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jairogloz/go-l/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s Service) Create(ctx context.Context, team *domain.Team) (err error) {
	now := time.Now().UTC()
	team.CreatedAt = primitive.DateTime(now.UnixNano() / int64(time.Millisecond))

	err = s.Repo.Insert(ctx, team)
	if err != nil {
		appErr := domain.ManageError(err, "")
		log.Println(appErr.Error())
		return fmt.Errorf("error creating team: %w", appErr)
	}

	return nil
}
