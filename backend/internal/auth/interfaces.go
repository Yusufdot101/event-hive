package auth

import (
	"github.com/Yusufdot101/eventhive/internal/token"
	"github.com/Yusufdot101/eventhive/internal/user"
)

type handler struct {
	userService  *user.UserService
	tokenService *token.TokenService
}

func NewHandler(userService *user.UserService, tokenService *token.TokenService) *handler {
	return &handler{
		userService:  userService,
		tokenService: tokenService,
	}
}
