package player

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-l/cmd/api/core"
	"github.com/jairogloz/go-l/internal/domain"
)

// GetPlayer godoc
// @Summary Get a player
// @Description Create a new player with the input payload
// @Tags players
// @Accept  json
// @Produce  json
// @Param PlayerIdParam body core.PlayerIdParam true "Create player"
// @Success 200 {object} map[string]interface{} "player: domain.Player"
// @Failure 400 {object} map[string]interface{} "error: string"
// @Router /players/:id [get]
func (h Handler) GetPlayer(c *gin.Context) {
	var playerIdParam core.PlayerIdParam
	if err := c.ShouldBindJSON(&playerIdParam); err != nil {
		core.RespondError(c, domain.NewAppError(domain.ErrCodeInvalidParams, err.Error()))
		return
	}

	player, err := h.PlayerService.Get(playerIdParam.Id)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.JSON(200, player)
}
