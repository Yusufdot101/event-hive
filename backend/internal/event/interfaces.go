package event

import (
	"time"

	"github.com/Yusufdot101/eventhive/internal/customerrors"
)

type event struct {
	id            string
	createdAt     time.Time `validate:"required"`
	startsAt      time.Time `validate:"required"`
	endsAt        time.Time `validate:"required"`
	lastUpdatedAt *time.Time
	creatorID     string  `validate:"required"`
	title         string  `validate:"required"`
	description   string  `validate:"required"`
	latitude      float64 `validate:"required,gte=-90,lte=90"`
	longitude     float64 `validate:"required,gte=-180,lte=180"`
	address       string  `validate:"required"`
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
	service *service
}

func newHandler(service *service) *handler {
	return &handler{
		service: service,
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
	if e.startsAt.Before(earliestDate) || e.endsAt.Before(e.startsAt) {
		return customerrors.ErrInvalidDates
	}
	return nil
}

func validateInfo(e *event) error {
	if len(e.title) < 2 || len(e.description) < 2 || len(e.address) < 2 {
		return customerrors.ErrInvalidInfo
	}
	return nil
}

func validateLocation(e *event) error {
	if (e.longitude > 180 || e.longitude < -180) || (e.latitude > 90 || e.latitude < -90) {
		return customerrors.ErrInvalidLocation
	}
	return nil
}
