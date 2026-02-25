package event

import (
	"errors"
	"net/http"

	"github.com/Yusufdot101/eventhive/internal/customerrors"
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

func (h *handler) getEvent(ctx *gin.Context) {
	eventID := ctx.Params.ByName("id")
	if eventID == "" {
		ctx.JSON(http.StatusNotFound, gin.H{"error": customerrors.ErrNoRecord})
		return
	}

	e, err := h.service.getByID(eventID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, customerrors.ErrNoRecord) {
			statusCode = http.StatusNotFound
		}
		if errors.Is(err, customerrors.ErrInvalidID) {
			statusCode = http.StatusBadRequest
		}
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"event": e,
	})
}

func (h *handler) GetEventAttendees(ctx *gin.Context) {
	eventID := ctx.Params.ByName("id")
	if eventID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": customerrors.ErrInvalidID.Error()})
		return
	}

	eventAttendees, err := h.attendanceService.GetEventAttendees(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userIDs := []string{}
	for _, attendee := range eventAttendees {
		userIDs = append(userIDs, attendee.UserID)
	}

	users, err := h.userService.GetUsersByIDs(userIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (h *handler) GetEventCreator(ctx *gin.Context) {
	eventID := ctx.Params.ByName("id")
	if eventID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": customerrors.ErrInvalidID.Error()})
		return
	}

	event, err := h.service.getByID(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.GetUserByID(event.CreatorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
