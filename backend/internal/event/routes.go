package event

import (
	"database/sql"
	"net/http"

	"github.com/Yusufdot101/eventhive/internal/attendance"
	"github.com/Yusufdot101/eventhive/internal/middleware"
	"github.com/Yusufdot101/eventhive/internal/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(DB *sql.DB, ctx *gin.RouterGroup) {
	repo := newRepository(DB)
	svc := newService(repo)

	userRepo := user.NewRepository(DB)
	userSvc := user.NewUserService(userRepo)

	attendanceRepo := attendance.NewRepository(DB)
	attendenceSvc := attendance.NewService(attendanceRepo)
	h := newHandler(svc, userSvc, attendenceSvc)

	attendenceHandler := attendance.NewHandler(attendenceSvc)

	ctx.Match([]string{http.MethodPost, http.MethodOptions}, "/events", middleware.Authenticate(h.create))
	ctx.Match([]string{http.MethodGet}, "/events", h.getEvents)
	ctx.Match([]string{http.MethodDelete}, "/events/:id", middleware.Authenticate(h.DeleteEvent))
	ctx.Match([]string{http.MethodGet}, "/events/:id", h.getEvent)

	ctx.Match([]string{http.MethodPost}, "/events/:id/attend", middleware.Authenticate(attendenceHandler.Create))
	ctx.Match([]string{http.MethodDelete}, "/events/:id/attend", middleware.Authenticate(attendenceHandler.Delete))

	ctx.Match([]string{http.MethodGet}, "/events/:id/attend", middleware.Authenticate(h.GetAttendingStatus))
	ctx.Match([]string{http.MethodGet}, "/events/:id/attendees", h.GetEventAttendees)

	ctx.Match([]string{http.MethodGet}, "/events/:id/creator", h.GetEventCreator)
}
