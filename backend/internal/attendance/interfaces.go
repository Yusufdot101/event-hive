package attendance

import (
	"time"

	"github.com/Yusufdot101/eventhive/internal/customerrors"
	"github.com/google/uuid"
)

type eventAttendee struct {
	id        string
	createdAt time.Time
	eventID   string
	userID    string
}

type service struct {
	repo *repository
}

func NewService(repo *repository) *service {
	return &service{
		repo: repo,
	}
}

type handler struct {
	svc *service
}

func NewHandler(svc *service) *handler {
	return &handler{
		svc: svc,
	}
}

func validateAttendee(ea *eventAttendee) error {
	err := uuid.Validate(ea.eventID)
	if err != nil {
		return customerrors.ErrInvalidID
	}

	err = uuid.Validate(ea.userID)
	if err != nil {
		return customerrors.ErrInvalidID
	}
	return nil
}
