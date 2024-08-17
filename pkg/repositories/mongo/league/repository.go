package league

import (
	"github.com/jairogloz/go-l/pkg/ports"
	"go.mongodb.org/mongo-driver/mongo"
)

// Make sure Repository implements ports.LeagueRepository
// at compile time
var _ ports.LeagueRepository = &Repository{}

// Repository is a struct that represents the repository for the league entity.
type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}
