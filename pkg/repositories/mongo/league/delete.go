package league

import (
	"context"
	"errors"
	"log/slog"

	"github.com/jairogloz/go-l/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrDeleteLeague = errors.New("error deleting league")
)

// Delete deletes a league by id from the database
func (r *Repository) Delete(id string) (err error) {

	leagueID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		slog.Error("converting id to object id: ", slog.Any("mongodb", err))
		return domain.ErrIncorrectID
	}

	collection := r.Client.Database("go-l").Collection("leagues")
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.M{"_id": leagueID})

	if err != nil {
		slog.Error("deleting league: ", slog.Any("mongodb", err))
		return ErrDeleteLeague
	}

	if deleteResult.DeletedCount == 0 {
		slog.Error("league not found")
		return domain.ErrNotFound
	}

	return nil
}
