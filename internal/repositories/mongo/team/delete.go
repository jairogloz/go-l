package team

import (
	"context"
	"errors"
	"log/slog"

	"github.com/jairogloz/go-l/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrDeleteTeam = errors.New("error deleting team")
)

// deletes one team by id from db
func (r Repository) Delete(ctx context.Context, id string) (err error) {

	teamID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		slog.Error("converting id to object id: ", slog.Any("mongodb", err))
		return domain.ErrIncorrectID
	}

	collection := r.Client.Database("go-l").Collection("teams")
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.M{"_id": teamID})
	if err != nil {
		slog.Error("deleting team: ", slog.Any("mongodb", err))
		return ErrDeleteTeam
	}

	if deleteResult.DeletedCount == 0 {
		slog.Error("player not found")
		return domain.ErrNotFound
	}

	return nil
}
