package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Yusufdot101/eventhive/internal/customerrors"
	"github.com/Yusufdot101/eventhive/internal/token"
	"github.com/gin-gonic/gin"
)

var signinRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *handler) signin(ctx *gin.Context) {
	if err := ctx.ShouldBind(&signinRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.userService.GetUserByEmailAndPassword(signinRequest.Email, signinRequest.Password)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, customerrors.ErrInvalidCredentials) {
			statusCode = http.StatusBadRequest
		}
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	refreshToken, err := h.tokenService.GenerateRefreshToken(token.RefreshToken, u.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := h.tokenService.GenerateJWT(token.AccessToken, u.ID)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("%v\n", err.Error()))
		return
	}

	ctx.SetCookie("refresh_token", refreshToken.TokenString, int(refreshToken.ExpiresAt.Unix()), "/auth", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"message":     "user logged in successfully",
		"accessToken": accessToken.TokenString,
	})
}
