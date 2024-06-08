package player

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-l/cmd/api/core"
)

// GetPlayer godoc
// @Summary Get a player
// @Description get a player with id param
// @Tags players
// @Accept  json
// @Produce  json
// @Param PlayerIdParam body core.PlayerIdParam true "Create player"
// @Success 200 {object} map[string]interface{} "player: domain.Player"
// @Failure 400 {object} map[string]interface{} "error: string"
// @Router /players/:id [get]
func (h Handler) GetPlayer(c *gin.Context) {
	playerIdParam := c.Param("id")
	player, err := h.PlayerService.Get(playerIdParam)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.JSON(200, player)
}
