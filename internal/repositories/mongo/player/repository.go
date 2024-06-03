package player

import (
	"github.com/jairogloz/go-l/internal/ports"
	"go.mongodb.org/mongo-driver/mongo"
)

// Make sure Repository implements ports.PlayerRepository
// at compile time
var _ ports.PlayerRepository = &Repository{}

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}
