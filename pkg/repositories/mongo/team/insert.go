package team

import (
	"context"
	"fmt"
	"github.com/jairogloz/go-l/pkg/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func (r *Repository) Insert(ctx context.Context, team *domain.Team) (err error) {
	team.ID = primitive.NewObjectID()

	_, err = r.Collection.InsertOne(ctx, team)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			log.Println("Duplicate key error")
			return fmt.Errorf("%w: error inserting team: %s",
				domain.ErrDuplicateKey, err.Error())
		}
		log.Println(err.Error())
		return fmt.Errorf("error inserting team: %w", err)
	}

	return nil
}
