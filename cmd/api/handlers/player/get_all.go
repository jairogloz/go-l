package player

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
