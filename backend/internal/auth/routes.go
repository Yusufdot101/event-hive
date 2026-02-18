package auth

import (
	"database/sql"
	"net/http"

	"github.com/Yusufdot101/eventhive/internal/token"
	"github.com/Yusufdot101/eventhive/internal/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(DB *sql.DB, ctx *gin.RouterGroup) {
	h := NewHandler(
		user.NewUserService(user.NewRepository(DB)),
		token.NewTokenService(token.NewRepository(DB)),
	)

	ctx.Match([]string{http.MethodPost, http.MethodOptions}, "/signup", h.Signup)
	ctx.Match([]string{http.MethodPut, http.MethodOptions}, "/signin", h.signin)
	ctx.Match([]string{http.MethodPut, http.MethodOptions}, "/logout", h.logout)
	ctx.Match([]string{http.MethodPut, http.MethodOptions}, "/refreshtoken", h.refreshToken)
}
