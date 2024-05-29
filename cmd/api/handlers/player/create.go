package player

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-l/internal/domain"
	"net/http"
)

func (h Handler) CreatePlayer(c *gin.Context) {
	// Traducir el request
	// int: validacion
	// consumir un servicio
	// traducir el response
	var playerCreateParams domain.Player
	if err := c.BindJSON(&playerCreateParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// ==================

	insertedId, err := h.PlayerService.Create(playerCreateParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oops!"})
	}

	// ==================

	c.JSON(200, gin.H{"player_id": insertedId})
}
