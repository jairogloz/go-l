package player

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllPlayers godoc
// @Summary Get all players
// @Description get all existing player
// @Tags players
// @Accept  json
// @Produce  json
// @Param PlayerIdParam body core.PlayerIdParam true "Create player"
// @Success 200 {object} map[string]interface{} "player: domain.Player"
// @Failure 400 {object} map[string]interface{} "error: string"
// @Router /players [get]
func (h Handler) GetAllPlayers(ctx *gin.Context) {
	players, err := h.PlayerService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if players == nil {
		ctx.Status(http.StatusNoContent)
	} else {
		ctx.JSON(http.StatusOK, players)
	}

}
