package event

import (
	"log"
	"net/http"
	"time"

	"github.com/Yusufdot101/eventhive/internal/middleware"
	"github.com/gin-gonic/gin"
)

var createRequest struct {
	// creatorID string, startsAt, endsAt time.Time, title, description string, latitude, longitude float64
	StartsAt    time.Time `json:"startsAt"`
	EndsAt      time.Time `json:"endsAt"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Address     string    `json:"address"`
}

func (h *handler) create(ctx *gin.Context) {
	userID := ctx.GetHeader(string(middleware.CtxUserKey))
	log.Println("here: ", userID)
	if err := ctx.ShouldBind(&createRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.newEvent(
		userID, createRequest.StartsAt, createRequest.EndsAt, createRequest.Title,
		createRequest.Description, createRequest.Latitude, createRequest.Longitude,
		createRequest.Address,
	)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "event created succesfully"})
}
