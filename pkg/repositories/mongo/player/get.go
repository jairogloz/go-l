package player

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/jairogloz/go-l/pkg/domain"
)

func (r *Repository) Get(id string) (player *domain.Player, err error) {
	playerID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	err = r.Collection.FindOne(context.Background(), bson.M{"_id": playerID}).Decode(&player)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	player.ID = playerID.Hex()
	return player, nil
}
