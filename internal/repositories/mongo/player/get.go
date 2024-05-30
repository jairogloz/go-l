package player

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/jairogloz/go-l/internal/domain"
)

type GetPlayerError struct {
	Err error
}

func (e GetPlayerError) Error() string {
	return fmt.Sprintf("error getting player: %s", e.Err.Error())
}

func (r *Repository) Get(id domain.Player) (player domain.Player, err error) {
	playerID, err := primitive.ObjectIDFromHex(id.ID)
	if err != nil {
		log.Println(err.Error())
		err = GetPlayerError{Err: err}
		return
	}

	collection := r.Client.Database("go-l").Collection("players")
	err = collection.FindOne(context.Background(), bson.M{"_id": playerID}).Decode(&player)
	if err != nil {
		log.Println(err.Error())
		err = GetPlayerError{Err: err}
		return
	}
	return
}
