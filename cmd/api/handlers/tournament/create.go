package tournament

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jairogloz/go-l/cmd/api/core"
	"github.com/jairogloz/go-l/pkg/domain"
)

// CreateTournament godoc
// @Summary Create a tournament
// @Description Create a tournament with the input payload
// @Tags tournaments
// @Accept  json
// @Produce  json
// @Param tournamentCreateParams body core.TournamentCreateParams true "Create tournament"
// @Success 201 {object} map[string]interface{} "tournament: domain.Tournament"
// @Failure 400 {object} map[string]interface{} "error: string"
// @Router /tournaments [post]
func (h Handler) CreateTournament(c *gin.Context) {
	var tournamentCreateParams core.TournamentCreateParams
	if err := c.ShouldBindJSON(&tournamentCreateParams); err != nil {
		core.RespondError(c, domain.NewAppError(domain.ErrCodeInvalidParams, err.Error()))
		return
	}

	tournament := &domain.Tournament{
		Name:        tournamentCreateParams.Name,
		Description: tournamentCreateParams.Description,
		URL:         tournamentCreateParams.URL,
	}
	err := h.TournamentService.Create(c, tournament)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.JSON(http.StatusCreated, tournament)
}
