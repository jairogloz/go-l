package tournament

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/jairogloz/go-l/pkg/ports"
)

// Make sure Repository implements ports.TournamentRepository
// at compile time
var _ ports.TournamentRepository = &Repository{}

// Repository is a struct that represents the repository for the tournament entity.
type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}
