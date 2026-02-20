package event

import (
	"database/sql"
	"net/http"

	"github.com/Yusufdot101/eventhive/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(DB *sql.DB, ctx *gin.RouterGroup) {
	repo := newRepository(DB)
	svc := newService(repo)
	h := newHandler(svc)

	ctx.Match([]string{http.MethodPost, http.MethodOptions}, "/events", middleware.Authenticate(h.create))
}
