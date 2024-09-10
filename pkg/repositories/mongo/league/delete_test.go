package league_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jairogloz/go-l/pkg/domain"
	"github.com/jairogloz/go-l/pkg/repositories/mongo"
	"github.com/jairogloz/go-l/pkg/repositories/mongo/league"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestRepository_DeleteLeague(t *testing.T) {
	// Generate a valid ObjectID
	validID := primitive.NewObjectID()

	// Skip this test if the short flag is provided
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	testTable := map[string]struct {
		leagueID      primitive.ObjectID
		setupFunc     func(repo *league.Repository) error
		assertionFunc func(subTest *testing.T, err error)
	}{
		"should delete league successfully": {
			leagueID: validID,
			setupFunc: func(repo *league.Repository) error {
				league := &domain.League{
					ID:   validID,
					Name: "Premier League",
				}
				return addLeagueToCollection(repo, league)
			},
			assertionFunc: func(subTest *testing.T, err error) {
				assert.Nil(subTest, err)
			},
		},
		"should return not found when deleting a non-existent league": {
			leagueID: primitive.NewObjectID(), 
			setupFunc: func(repo *league.Repository) error {
				return nil
			},
			assertionFunc: func(subTest *testing.T, err error) {
				assert.ErrorContains(subTest, err, domain.ErrNotFound.Error())
			},
		},
		"should return error when given an invalid league ID": {
			leagueID: primitive.NilObjectID, // Using a nil ObjectID to simulate invalid
			setupFunc: func(repo *league.Repository) error {
				return nil
			},
			assertionFunc: func(subTest *testing.T, err error) {
				assert.ErrorContains(subTest, err, domain.ErrNotFound.Error())
			},
		},
	}

	db, err := startMongoDB(context.Background())
	if err != nil {
		t.Fatalf("error starting mongodb container: %v", err)
	}
	defer func() {
		terminateErr := db.Container.Terminate(context.Background())
		if terminateErr != nil {
			t.Errorf("error terminating mongodb container: %v", terminateErr)
		}
	}()

	repo, err := setupLeagueRepository(db)
	if err != nil {
		t.Fatalf("error setting up repository: %v", err)
	}

	for testName, test := range testTable {
		t.Run(testName, func(subTest *testing.T) {
			if test.setupFunc != nil {
				if err := test.setupFunc(repo); err != nil {
					subTest.Fatalf("error in setup: %v", err)
				}
			}
			err := repo.Delete(test.leagueID.Hex()) // Convert ObjectID to string
			test.assertionFunc(subTest, err)
		})
	}
}

// setupLeagueRepository will return a league repository instance or an error
func setupLeagueRepository(db *mongodb.MongoDBContainer) (*league.Repository, error) {
	connString, err := db.ConnectionString(context.Background())
	if err != nil {
		return nil, err
	}
	client, err := mongo.ConnectClient(connString)
	if err != nil {
		return nil, err
	}
	leagueRepo := &league.Repository{
		Client:     client,
		Collection: client.Database("go-l").Collection("leagues"),
	}
	if err = leagueRepo.CreateIndexes(); err != nil {
		return nil, err
	}
	return leagueRepo, nil
}

// addLeagueToCollection inserts a league with a specific ObjectID
func addLeagueToCollection(repo *league.Repository, l *domain.League) error {
	_, err := repo.Collection.InsertOne(context.Background(), l)
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
