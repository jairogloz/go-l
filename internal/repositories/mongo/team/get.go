package team

import (
	"context"
	"fmt"
	"log"
	"github.com/jairogloz/go-l/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r Repository) Get(ctx context.Context, id string) (team domain.Team, err error) {
	teamID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err.Error())
		return team, fmt.Errorf("%w: error getting team: %s", domain.ErrIncorrectID, err)
	}

	collection := r.Client.Database("go-l").Collection("teams")
	err = collection.FindOne(ctx, bson.M{"_id": teamID}).Decode(&team)
	if err != nil {
		log.Println(err.Error())
		return team, fmt.Errorf("%w: error getting team: %s", domain.ErrNotFound, err)
	}

	team.ID = teamID.Hex()
	return
}
