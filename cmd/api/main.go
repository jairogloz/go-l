package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/jairogloz/go-l/cmd/api/handlers/player"
	"github.com/jairogloz/go-l/cmd/api/handlers/tournament"
	"github.com/jairogloz/go-l/pkg/repositories/mongo"
	playerMongo "github.com/jairogloz/go-l/pkg/repositories/mongo/player"
	tournamentMongo "github.com/jairogloz/go-l/pkg/repositories/mongo/tournament"
	playerService "github.com/jairogloz/go-l/pkg/services/player"
	tournamentService "github.com/jairogloz/go-l/pkg/services/tournament"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ginEngine := gin.Default()

	client, err := mongo.ConnectClient(os.Getenv("MONGO_URI"))
	if err != nil {
		log.Fatal(err.Error())
	}

	database := client.Database("go-l")

	playerRepo := &playerMongo.Repository{
		Client:     client,
		Collection: database.Collection("players"),
	}

	tournamentRepo := &tournamentMongo.Repository{
		Client:     client,
		Collection: database.Collection("tournaments"),
	}

	playerSrv := &playerService.Service{
		Repo: playerRepo,
	}

	tournamentSrv := &tournamentService.Service{
		Repo: tournamentRepo,
	}

	playerHandler := &player.Handler{
		PlayerService: playerSrv,
	}

	tournamentHandler := &tournament.Handler{
		TournamentService: tournamentSrv,
	}

	ginEngine.GET("/players/:id", playerHandler.GetPlayer)
	ginEngine.POST("/players", playerHandler.CreatePlayer)
	ginEngine.DELETE("/players/:id", playerHandler.DeletePlayer)

	ginEngine.POST("/tournaments", tournamentHandler.CreateTournament)

	log.Fatalln(ginEngine.Run(":8001"))

}
