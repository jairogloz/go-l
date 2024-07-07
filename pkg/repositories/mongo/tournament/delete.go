package tournament

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/jairogloz/go-l/pkg/domain"
)

// Delete deletes a tournament by id from db
func (r *Repository) Delete(ctx context.Context, id string) (err error) {
	teamID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.ErrIncorrectID
	}

	deleteResult, err := r.Collection.DeleteOne(ctx, bson.M{"_id": teamID})
	if err != nil {
		return fmt.Errorf("error deleting tournament: %s", err.Error())
	}

	if deleteResult.DeletedCount == 0 {
		return domain.ErrNotFound
	}

	return nil
}
