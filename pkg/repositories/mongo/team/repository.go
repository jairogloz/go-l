package team

import (
	"github.com/jairogloz/go-l/pkg/ports"
	"go.mongodb.org/mongo-driver/mongo"
)

// Make sure Repository implements ports.TeamRepository
// at compile time
var _ ports.TeamRepository = &Repository{}

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}
