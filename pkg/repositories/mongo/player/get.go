package player

import (
	"context"
	"errors"
	"fmt"
	"github.com/jairogloz/go-l/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get retrieves a player from the MongoDB collection by their ID.
// It first converts the provided string ID to a MongoDB ObjectID.
// If the conversion fails (e.g., if the ID is not a valid hex string), it returns an error.
// It then attempts to find a document in the MongoDB collection with the converted ObjectID.
// If it finds a document, it decodes it into a domain.Player object.
// If it doesn't find a document, it returns a domain-specific not found error.
// If there is a timeout error when accessing the MongoDB collection, it returns a domain-specific timeout error.
// For any other errors, it returns the error as is.
// If it successfully finds and decodes a document, it sets the ID of the domain.Player object to the hex string representation of the ObjectID and returns the domain.Player object.
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
		var commandErr *mongo.CommandError
		if errors.As(err, &commandErr) && commandErr.HasErrorLabel("NetworkTimeout") {
			return nil, domain.ErrTimeout
		}
		return nil, err
	}

	player.ID = playerID.Hex()
	return player, nil
}
