package attendance

import (
	"time"

	"github.com/Yusufdot101/eventhive/internal/customerrors"
	"github.com/google/uuid"
)

type eventAttendee struct {
	CreatedAt time.Time
	EventID   string
	UserID    string
}

type Service struct {
	repo *repository
}

func NewService(repo *repository) *Service {
	return &Service{
		repo: repo,
	}
}

type handler struct {
	svc *Service
}

func NewHandler(svc *Service) *handler {
	return &handler{
		svc: svc,
	}
}

func validateAttendee(ea *eventAttendee) error {
	err := uuid.Validate(ea.EventID)
	if err != nil {
		return customerrors.ErrInvalidID
	}

	err = uuid.Validate(ea.UserID)
	if err != nil {
		return customerrors.ErrInvalidID
	}
	return nil
}
