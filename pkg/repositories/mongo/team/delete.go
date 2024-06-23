package team

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jairogloz/go-l/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete deletes a team by id from db
func (r Repository) Delete(ctx context.Context, id string) (err error) {

	teamID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		slog.Error("converting id to object id: ", slog.Any("mongodb", err))
		return domain.ErrIncorrectID
	}

	deleteResult, err := r.Collection.DeleteOne(ctx, bson.M{"_id": teamID})
	if err != nil {
		slog.Error("deleting team: ", slog.Any("mongodb", err))
		return fmt.Errorf("error deleting team: %s", err.Error())
	}

	if deleteResult.DeletedCount == 0 {
		slog.Error("player not found")
		return domain.ErrNotFound
	}

	return nil
}
