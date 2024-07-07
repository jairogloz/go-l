package player

import (
	"context"
	"fmt"

	"github.com/jairogloz/go-l/pkg/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Make sure Repository implements ports.PlayerRepository
// at compile time
var _ ports.PlayerRepository = &Repository{}

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

// CreateIndexes creates the secondary indexes for the collection
func (r *Repository) CreateIndexes() error {
	// Create the team_id index
	// This index is used to search for players by team_id
	// The index is not unique because a team can have multiple players
	_, err := r.Collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "team_info.team_id", Value: 1}},
		Options: options.Index().SetUnique(false),
	})
	if err != nil {
		return fmt.Errorf("error creating team_id index: %w", err)
	}

	return nil
}
