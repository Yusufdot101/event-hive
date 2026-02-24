package event

import (
	"time"

	"github.com/Yusufdot101/eventhive/internal/attendance"
	"github.com/Yusufdot101/eventhive/internal/customerrors"
	"github.com/Yusufdot101/eventhive/internal/user"
)

type event struct {
	ID            string
	CreatedAt     time.Time `validate:"required"`
	StartsAt      time.Time `validate:"required"`
	EndsAt        time.Time `validate:"required"`
	LastUpdatedAt *time.Time
	CreatorID     string  `validate:"required"`
	Title         string  `validate:"required"`
	Description   string  `validate:"required"`
	Latitude      float64 `validate:"required,gte=-90,lte=90"`
	Longitude     float64 `validate:"required,gte=-180,lte=180"`
	Address       string  `validate:"required"`
}

type service struct {
	repo *repository
}

func newService(repo *repository) *service {
	return &service{
		repo: repo,
	}
}

type handler struct {
	service           *service
	userService       *user.UserService
	attendanceService *attendance.Service
}

func newHandler(service *service, userSvc *user.UserService, attendanceSvc *attendance.Service) *handler {
	return &handler{
		service:           service,
		userService:       userSvc,
		attendanceService: attendanceSvc,
	}
}

func validateEvent(e *event) error {
	// validate dates
	err := validateDates(e)
	if err != nil {
		return err
	}

	// validate event info
	err = validateInfo(e)
	if err != nil {
		return err
	}

	// valdiate location
	err = validateLocation(e)
	if err != nil {
		return err
	}
	return nil
}

func validateDates(e *event) error {
	earliestDate := time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC)
	if e.StartsAt.Before(earliestDate) || e.EndsAt.Before(e.StartsAt) {
		return customerrors.ErrInvalidDates
	}
	return nil
}

func validateInfo(e *event) error {
	if len(e.Title) < 2 || len(e.Description) < 2 || len(e.Address) < 2 {
		return customerrors.ErrInvalidInfo
	}
	return nil
}

func validateLocation(e *event) error {
	if (e.Longitude > 180 || e.Longitude < -180) || (e.Latitude > 90 || e.Latitude < -90) {
		return customerrors.ErrInvalidLocation
	}
	return nil
}
