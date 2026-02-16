package auth

import (
	"errors"
	"net/http"

	"github.com/Yusufdot101/eventhive/internal/customerrors"
	"github.com/Yusufdot101/eventhive/internal/token"
	"github.com/gin-gonic/gin"
)

func (h *handler) refreshToken(ctx *gin.Context) {
	cookie, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tk, err := h.tokenService.GetTokenByStringAndUse(cookie, token.RefreshToken)
	if err != nil {
		if errors.Is(err, customerrors.ErrInvalidRefreshToken) {
			// delete the cookie
			ctx.SetCookie("refresh_token", cookie, -1, "/auth", "", false, true)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := h.tokenService.GenerateJWT(token.AccessToken, tk.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":     "access token refreshed successfully",
		"accessToken": accessToken.TokenString,
	})
}
