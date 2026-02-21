package event

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) getEvents(ctx *gin.Context) {
	events, err := h.service.getMany()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}
