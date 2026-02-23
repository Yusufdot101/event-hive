package attendance

import (
	"errors"

	"github.com/Yusufdot101/eventhive/internal/customerrors"
)

func (svc *service) NewAttendance(eventID, userID string) error {
	ea := &eventAttendee{
		eventID: eventID,
		userID:  userID,
	}

	err := validateAttendee(ea)
	if err != nil {
		return err
	}
	return svc.repo.insert(ea)
}

func (svc *service) RemoveAttendance(eventID, userID string) error {
	ea := &eventAttendee{
		eventID: eventID,
		userID:  userID,
	}

	err := validateAttendee(ea)
	if err != nil {
		return err
	}
	return svc.repo.delete(ea)
}

func (svc *service) UserIsAttendingEvent(eventID, userID string) (bool, error) {
	ea := &eventAttendee{
		eventID: eventID,
		userID:  userID,
	}

	err := validateAttendee(ea)
	if err != nil {
		return false, err
	}
	_, err = svc.repo.get(ea)
	if err != nil && !errors.Is(err, customerrors.ErrNoRecord) {
		return false, err
	}

	if err != nil && errors.Is(err, customerrors.ErrNoRecord) {
		return false, nil
	}

	return true, nil
}
