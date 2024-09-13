package league

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jairogloz/go-l/cmd/api/core"
)

// DeleteLeague godoc
// @Summary Delete a league
// @Description Delete a league by ID
// @Tags leagues
// @Accept  json
// @Produce  json
// @Param id path string true "League ID"
// @Success 204 {object} map[string]interface{} "No content"
// @Failure 400 {object} map[string]interface{} "error: string"
// @Router /leagues/{id} [delete]
func (h Handler) DeleteLeague(c *gin.Context) {
	id := c.Param("id")

	err := h.LeagueService.Delete(id)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
