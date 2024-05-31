package player

import (
	"context"
	"errors"
	"log/slog"

	"github.com/jairogloz/go-l/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrPlayerNotFound = errors.New("player not found")
	ErrDeletePlayer   = errors.New("error deleting player")
	ErrIncorrectHexID = errors.New("incorrect mongo Hex format")
)

// Used for delete one player by id from the database
func (r *Repository) Delete(id domain.Player) (err error) {

	playerID, err := primitive.ObjectIDFromHex(id.ID)
	if err != nil {
		slog.Error("converting id to object id: ", slog.Any("mongodb", err))
		return ErrIncorrectHexID
	}

	collection := r.Client.Database("go-l").Collection("players")
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.M{"_id": playerID})

	if err != nil {
		slog.Error("deleting player: ", slog.Any("mongodb", err))
		return ErrDeletePlayer
	}

	if deleteResult.DeletedCount == 0 {
		slog.Error("player not found")
		return ErrPlayerNotFound
	}

	return nil
}
