package middleware

import (
	"net/http"
	"strings"

	"github.com/Yusufdot101/eventhive/internal/customerrors"
	"github.com/Yusufdot101/eventhive/internal/token"
	"github.com/gin-gonic/gin"
)

type ContextKey string

const CtxUserKey ContextKey = "userID"

func Authenticate(next gin.HandlerFunc) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": customerrors.ErrInvalidAccessToken.Error()})
			return
		}

		headParts := strings.Split(authHeader, " ")
		if len(headParts) != 2 || headParts[0] != "Bearer" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": customerrors.ErrInvalidAccessToken.Error()})
			return
		}

		tokenString := headParts[1]

		tk, err := token.ValidateJWT(tokenString)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": customerrors.ErrInvalidAccessToken.Error()})
			return
		}

		userID, err := tk.Claims.GetSubject()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": customerrors.ErrInvalidAccessToken.Error()})
			return
		}

		ctx.Request.Header.Set(string(CtxUserKey), userID)
		next(ctx)
	}
	return fn
}
