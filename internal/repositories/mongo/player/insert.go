package player

import (
	"context"
	"fmt"
	"github.com/jairogloz/go-l/internal/domain"
	"log"
)

func (r *Repository) Insert(player domain.Player) (id interface{}, err error) {

	// ===============
	collection := r.Client.Database("go-l").Collection("players")
	insertResult, err := collection.InsertOne(context.Background(), player)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error inserting player: %w", err)
	}

	return insertResult.InsertedID, nil
}
