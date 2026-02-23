package event

import (
	"database/sql"
	"net/http"

	"github.com/Yusufdot101/eventhive/internal/attendance"
	"github.com/Yusufdot101/eventhive/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(DB *sql.DB, ctx *gin.RouterGroup) {
	repo := newRepository(DB)
	svc := newService(repo)
	h := newHandler(svc)

	attendenceRepo := attendance.NewRepository(DB)
	attendenceService := attendance.NewService(attendenceRepo)
	attendenceHandler := attendance.NewHandler(attendenceService)

	ctx.Match([]string{http.MethodPost, http.MethodOptions}, "/events", middleware.Authenticate(h.create))
	ctx.Match([]string{http.MethodGet}, "/events", h.getEvents)
	ctx.Match([]string{http.MethodGet}, "/events/:id", h.getEvent)
	ctx.Match([]string{http.MethodPost}, "/events/:id/attend", middleware.Authenticate(attendenceHandler.Create))
	ctx.Match([]string{http.MethodDelete}, "/events/:id/attend", middleware.Authenticate(attendenceHandler.Delete))
	ctx.Match([]string{http.MethodGet}, "/events/:id/attend", middleware.Authenticate(attendenceHandler.GetAttendingStatus))
}
