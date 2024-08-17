package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-l/cmd/api/handlers/league"
	"github.com/joho/godotenv"

	"github.com/jairogloz/go-l/cmd/api/handlers/league"
	"github.com/jairogloz/go-l/cmd/api/handlers/player"
	"github.com/jairogloz/go-l/cmd/api/handlers/tournament"
	"github.com/jairogloz/go-l/pkg/repositories/mongo"
	leagueMongo "github.com/jairogloz/go-l/pkg/repositories/mongo/league"
	playerMongo "github.com/jairogloz/go-l/pkg/repositories/mongo/player"
	tournamentMongo "github.com/jairogloz/go-l/pkg/repositories/mongo/tournament"
	leagueService "github.com/jairogloz/go-l/pkg/services/league"
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

	leagueRepo := &leagueMongo.Repository{
		Client:     client,
		Collection: database.Collection("leagues"),
	}

	err = playerRepo.CreateIndexes()
	if err != nil {
		log.Fatal(err.Error())
	}

	playerSrv := &playerService.Service{
		Repo: playerRepo,
	}

	tournamentSrv := &tournamentService.Service{
		Repo: tournamentRepo,
	}

	leagueSrv := &leagueService.Service{
		Repo: leagueRepo,
	}

	playerHandler := &player.Handler{
		PlayerService: playerSrv,
	}

	leagueRepo := &leagueMongo.Repository{
		Client:     client,
		Collection: client.Database("go-l").Collection("leagues"),
	}

	leagueSrv := &leagueService.Service{
		Repo: leagueRepo,
	}

	leagueHandler := &league.Handler{
		LeagueService: leagueSrv,
	}

	tournamentHandler := &tournament.Handler{
		TournamentService: tournamentSrv,
	}

	leagueHandler := &league.Handler{
		LeagueService: leagueSrv,
	}

	ginEngine.GET("/players/:id", playerHandler.GetPlayer)
	ginEngine.POST("/players", playerHandler.CreatePlayer)
	ginEngine.DELETE("/players/:id", playerHandler.DeletePlayer)

	ginEngine.GET("/league/:id", leagueHandler.GetLeague)
	ginEngine.POST("/tournaments", tournamentHandler.CreateTournament)
	ginEngine.DELETE("/tournaments/:id", tournamentHandler.DeleteTournament)

	ginEngine.POST("/leagues", leagueHandler.CreateLeague)

	log.Fatalln(ginEngine.Run(":8001"))

}
