package league

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-l/cmd/api/core"
)

// GetLeague godoc
// @Summary Get a league
// @Description get a league with id param
// @Tags league
// @Accept  json
// @Produce  json
// @Param LeagueIdParam body core.LeagueIdParam true "Create league"
// @Success 200 {object} "league: domain.League"
// @Failure 400 {object} map[string]interface{} "error: string"
// @Router /league/:id [get]
func (h Handler) GetLeague(c *gin.Context) {
	leagueIdParam := c.Param("id")
	player, err := h.LeagueService.Get(leagueIdParam)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.JSON(200, player)
}
