package league

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jairogloz/go-l/cmd/api/core"
	"github.com/jairogloz/go-l/pkg/domain"
)

// CreateLeague godoc
// @Summary Create a league
// @Description Create a league with the input payload
// @Tags leagues
// @Accept  json
// @Produce  json
// @Param leagueCreateParams body core.LeagueCreateParams true "Create league"
// @Success 201 {object} map[string]interface{} "league: domain.League"
// @Failure 400 {object} map[string]interface{} "error: string"
// @Failure 500 {object} map[string]interface{} "error: string"
// @Router /leagues [post]
func (h *Handler) CreateLeague(c *gin.Context) {
	var leagueCreateParams core.LeagueCreateParams
	if err := c.ShouldBindJSON(&leagueCreateParams); err != nil {
		core.RespondError(c, domain.NewAppError(domain.ErrCodeInvalidParams, err.Error()))
		return
	}

	league := &domain.League{
		Name:        leagueCreateParams.Name,
		Description: leagueCreateParams.Description,
	}
	err := h.LeagueService.Create(c, league)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.JSON(http.StatusCreated, league)
}
