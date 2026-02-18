package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Yusufdot101/eventhive/internal/config"
	"github.com/Yusufdot101/eventhive/internal/customerrors"
	"github.com/Yusufdot101/eventhive/internal/token"
	"github.com/gin-gonic/gin"
)

var signupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *handler) Signup(ctx *gin.Context) {
	if err := ctx.ShouldBind(&signupRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.userService.RegisterUser(signupRequest.Name, signupRequest.Email, signupRequest.Password)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, customerrors.ErrDuplicateEmail) {
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

	ctx.SetCookie("refresh_token", refreshToken.TokenString, int(refreshToken.ExpiresAt.Unix()), "/auth", "", config.SecureCookie, true)

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, gin.H{
		"message":     "user created successfully",
		"accessToken": accessToken.TokenString,
	})
}
