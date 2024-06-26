package player

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-l/cmd/api/core"
	"github.com/jairogloz/go-l/pkg/domain"
)

// CreatePlayer godoc
// @Summary Create a new player
// @Description Create a new player with the input payload
// @Tags players
// @Accept  json
// @Produce  json
// @Param playerCreateParams body core.PlayerCreateParams true "Create player"
// @Success 200 {object} map[string]interface{} "player_id: string"
// @Failure 400 {object} map[string]interface{} "error: string"
// @Router /players [post]
func (h Handler) CreatePlayer(c *gin.Context) {
	var playerCreateParams core.PlayerCreateParams
	if err := c.ShouldBindJSON(&playerCreateParams); err != nil {
		core.RespondError(c, domain.NewAppError(domain.ErrCodeInvalidParams, err.Error()))
		return
	}

	player := &domain.Player{
		ContactInfo: playerCreateParams.ContactInfo,
		DateOfBirth: playerCreateParams.DateOfBirth,
		FirstName:   playerCreateParams.FirstName,
		LastName:    playerCreateParams.LastName,
		TeamInfo:    playerCreateParams.TeamInfo,
	}
	err := h.PlayerService.Create(player)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.JSON(200, player)
}
