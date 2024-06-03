package player

import (
	"context"
	"fmt"
	"github.com/jairogloz/go-l/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func (r *Repository) Insert(player *domain.Player) (err error) {

	player.ID = primitive.NewObjectID()

	_, err = r.Collection.InsertOne(context.Background(), player)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			log.Println("Duplicate key error")
			return fmt.Errorf("%w: error inserting player: %s",
				domain.ErrDuplicateKey, err.Error())
		}
		log.Println(err.Error())
		return fmt.Errorf("error inserting player: %w", err)
	}

	return nil
}
