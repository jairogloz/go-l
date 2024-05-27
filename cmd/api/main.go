package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-l/cmd/api/core"
	"log"
)

func main() {

	server := core.Server{GinEngine: gin.Default()}

	server.GinEngine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong!"})
	})

	log.Fatalln(server.GinEngine.Run(":8001"))

}
