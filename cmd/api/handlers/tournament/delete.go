package tournament

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jairogloz/go-l/cmd/api/core"
)

// DeleteTournament godoc
// @Summary Delete a tournament
// @Description Delete a tournament by ID
// @Tags tournaments
// @Accept  json
// @Produce  json
// @Param id path string true "Tournament ID"
// @Success 204 {object} map[string]interface{} "No content"
// @Failure 400 {object} map[string]interface{} "error: string"
// @Router /tournaments/{id} [delete]
func (h Handler) DeleteTournament(c *gin.Context) {
	id := c.Param("id")

	err := h.TournamentService.Delete(c, id)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
