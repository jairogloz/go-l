package league

import (
	"github.com/jairogloz/go-l/pkg/ports"
	"go.mongodb.org/mongo-driver/mongo"
)

// Make sure Repository implements ports.PlayerRepository
// at compile time
var _ ports.LeagueRepository = &Repository{}

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}
