package player

import (
	"context"
	"fmt"
	"github.com/jairogloz/go-l/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetAll retrieves all players from the MongoDB collection.
func (r *Repository) GetAll() ([]*domain.Player, error) {
	cursor, err := r.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error getting players: %w", err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			fmt.Printf("error closing cursor: %v", err)
		}
	}(cursor, context.Background())

	var players []*domain.Player
	for cursor.Next(context.Background()) {
		var player domain.Player
		err := cursor.Decode(&player)
		if err != nil {
			return nil, fmt.Errorf("error decoding player: %w", err)
		}
		players = append(players, &player)
	}

	return players, nil
}
