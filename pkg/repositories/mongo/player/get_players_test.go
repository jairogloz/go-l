package player_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jairogloz/go-l/pkg/domain"
	"github.com/jairogloz/go-l/pkg/repositories/mongo"
	"github.com/jairogloz/go-l/pkg/repositories/mongo/player"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
)

func TestService_GetPlayersByTeamID(t *testing.T) {
	// Skip this test if the short flag is provided
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	testTable := map[string]struct {
		teamID        string
		assertionFunc func(subTest *testing.T, p []*domain.Player, err error)
	}{
		"should return players by team id": {
			teamID: "team1",
			assertionFunc: func(subTest *testing.T, p []*domain.Player, err error) {
				assert.Nil(subTest, err)
				assert.Equal(subTest, 1, len(p))
			},
		},
		"should return empty players list when team id not found": {
			teamID: "team2",
			assertionFunc: func(subTest *testing.T, p []*domain.Player, err error) {
				assert.ErrorContains(subTest, err, domain.ErrNotFound.Error())
				assert.Equal(subTest, 0, len(p))
			},
		},
	}

	db, err := startMongoDB(context.Background())
	if err != nil {
		t.Fatalf("error starting mongodb container: %v", err)
	}
	defer db.Container.Terminate(context.Background()) // nolint: errcheck
	repo, err := setupRepository(db)
	if err != nil {
		t.Fatalf("error setting up repository: %v", err)
	}
	if err := addPlayerToCollection(repo, &domain.Player{TeamInfo: &domain.TeamInfo{TeamID: "team1"}}); err != nil {
		t.Fatalf("error adding player to collection: %v", err)
	}
	for testName, test := range testTable {
		t.Run(testName, func(subTest *testing.T) {
			players, err := repo.GetPlayersByTeamID(test.teamID)
			test.assertionFunc(subTest, players, err)
		})
	}
}

// setupRepository will return a player repository instance or an error
// It will create an index in the collection for testing purposes
// It will also create a player document in the collection for testing purposes
func setupRepository(db *mongodb.MongoDBContainer) (*player.Repository, error) {
	connString, err := db.ConnectionString(context.Background())
	if err != nil {
		return nil, err
	}
	client, err := mongo.ConnectClient(connString)
	if err != nil {
		return nil, err

	}
	playerRepo := &player.Repository{
		Client:     client,
		Collection: client.Database("go-l").Collection("players"),
	}
	if err = playerRepo.CreateIndexes(); err != nil {
		return nil, err
	}
	return playerRepo, nil
}

func addPlayerToCollection(repo *player.Repository, p *domain.Player) error {
	_, err := repo.Collection.InsertOne(context.Background(), p)
	if err != nil {
		return err
	}
	return nil
}

// startMongoDB will return a mongodb testcontainer instance or an error
func startMongoDB(ctx context.Context) (*mongodb.MongoDBContainer, error) {
	mongodbContainer, err := mongodb.RunContainer(ctx)
	if err != nil {
		return nil, fmt.Errorf("error running mongodb container: %w", err)
	}
	return mongodbContainer, nil
}
