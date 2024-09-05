package league

import (
	"context"
	"fmt"
	
	"github.com/jairogloz/go-l/pkg/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Make sure Repository implements ports.LeagueRepository
// at compile time
var _ ports.LeagueRepository = &Repository{}

// Repository is a struct that represents the repository for the league entity.
type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

// CreateIndexes creates the secondary indexes for the collection
func (r *Repository) CreateIndexes() error {
	// Create the league_id index
	// This index is used to search for leagues by league_id
	// The index is not unique because a league can have multiple records
	_, err := r.Collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "league_info.league_id", Value: 1}},
		Options: options.Index().SetUnique(false),
	})
	if err != nil {
		return fmt.Errorf("error creating league_id index: %w", err)
	}

	// Add any additional indexes if needed

	return nil
}
