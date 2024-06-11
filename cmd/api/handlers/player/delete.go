package player

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jairogloz/go-l/cmd/api/core"
)

// DeletePlayer godoc
// @Summary Delete a player
// @Description Delete a player by ID
// @Tags players
// @Accept  json
// @Produce  json
// @Param id path string true "Player ID"
// @Success 204 {object} map[string]interface{} "No content"
// @Failure 400 {object} map[string]interface{} "error: string"
// @Router /players/{id} [delete]
func (h Handler) DeletePlayer(c *gin.Context) {
	id := c.Param("id")

	err := h.PlayerService.Delete(id)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
