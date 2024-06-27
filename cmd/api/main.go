package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-l/cmd/api/handlers/league"
	"github.com/jairogloz/go-l/cmd/api/handlers/player"
	"github.com/jairogloz/go-l/pkg/repositories/mongo"
	leagueMongo "github.com/jairogloz/go-l/pkg/repositories/mongo/league"
	playerMongo "github.com/jairogloz/go-l/pkg/repositories/mongo/player"
	leagueService "github.com/jairogloz/go-l/pkg/services/league"
	playerService "github.com/jairogloz/go-l/pkg/services/player"
	"github.com/joho/godotenv"
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

	playerRepo := &playerMongo.Repository{
		Client:     client,
		Collection: client.Database("go-l").Collection("players"),
	}

	playerSrv := &playerService.Service{
		Repo: playerRepo,
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

	ginEngine.GET("/players/:id", playerHandler.GetPlayer)
	ginEngine.POST("/players", playerHandler.CreatePlayer)
	ginEngine.DELETE("/players/:id", playerHandler.DeletePlayer)


	ginEngine.GET("/league/:id", leagueHandler.GetLeague)

	log.Fatalln(ginEngine.Run(":8001"))

}
