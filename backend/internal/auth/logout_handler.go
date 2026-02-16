package auth

import (
	"net/http"

	"github.com/Yusufdot101/eventhive/internal/token"
	"github.com/gin-gonic/gin"
)

func (h *handler) logout(ctx *gin.Context) {
	// get the refresh token
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "refresh_token missing"})
		return
	}

	// delete the token
	err = h.tokenService.DeleteTokenByStringAndUse(refreshToken, token.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// send message informing client
	ctx.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
