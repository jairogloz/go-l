package player

import (
	"context"
	"errors"
	"log/slog"

	"github.com/jairogloz/go-l/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	errPlayerNotFound = errors.New("player not found")
	errDeletePlayer   = errors.New("error deleting player")
)

// Used for delete one player by id from the database
func (r *Repository) Delete(id domain.Player) (err error) {

	collection := r.Client.Database("go-l").Collection("players")
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		slog.Error("deleting player: ", slog.Any("mongodb", err))
		return errDeletePlayer
	}

	if deleteResult.DeletedCount == 0 {
		slog.Error("player not found")
		return errPlayerNotFound
	}

	return nil
}
