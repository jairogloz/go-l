package player

import (
	"context"
	"errors"
	"fmt"

	"github.com/jairogloz/go-l/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetPlayersByTeamID retrieves all players from the MongoDB collection that belong to a team with the provided team ID.
func (r *Repository) GetPlayersByTeamID(teamID string) ([]*domain.Player, error) {
	cursor, err := r.Collection.Find(context.Background(), bson.M{"team_info.team_id": teamID})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("error getting players: %w", err)
	}
	defer cursor.Close(context.Background())

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
