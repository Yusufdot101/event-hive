package attendance

import (
	"net/http"

	"github.com/Yusufdot101/eventhive/internal/customerrors"
	"github.com/Yusufdot101/eventhive/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (h *handler) Delete(ctx *gin.Context) {
	userID := ctx.GetHeader(string(middleware.CtxUserKey))
	if userID == "" {
		panic("userID should be in the header")
	}

	eventID := ctx.Params.ByName("id")
	if eventID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": customerrors.ErrInvalidID.Error()})
		return
	}

	err := h.svc.RemoveAttendance(eventID, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "attendance removed successfully"})
}
