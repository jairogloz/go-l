package player

import (
	"context"
	"errors"
	"log/slog"

	"github.com/jairogloz/go-l/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrDeletePlayer = errors.New("error deleting player")
)

// Used for delete one player by id from the database
func (r *Repository) Delete(id string) (err error) {

	playerID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		slog.Error("converting id to object id: ", slog.Any("mongodb", err))
		return domain.ErrIncorrectID
	}

	collection := r.Client.Database("go-l").Collection("players")
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.M{"_id": playerID})

	if err != nil {
		slog.Error("deleting player: ", slog.Any("mongodb", err))
		return ErrDeletePlayer
	}

	if deleteResult.DeletedCount == 0 {
		slog.Error("player not found")
		return domain.ErrNotFound
	}

	return nil
}
