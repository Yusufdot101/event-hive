package attendance

import (
	"errors"

	"github.com/Yusufdot101/eventhive/internal/customerrors"
)

func (svc *Service) NewAttendance(eventID, userID string) error {
	ea := &eventAttendee{
		EventID: eventID,
		UserID:  userID,
	}

	err := validateAttendee(ea)
	if err != nil {
		return err
	}
	return svc.repo.insert(ea)
}

func (svc *Service) RemoveAttendance(eventID, userID string) error {
	ea := &eventAttendee{
		EventID: eventID,
		UserID:  userID,
	}

	err := validateAttendee(ea)
	if err != nil {
		return err
	}
	return svc.repo.delete(ea)
}

func (svc *Service) UserIsAttendingEvent(eventID, userID string) (bool, error) {
	ea := &eventAttendee{
		EventID: eventID,
		UserID:  userID,
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

func (svc *Service) GetEventAttendees(eventID string) ([]*eventAttendee, error) {
	return svc.repo.getManyByEventID(eventID)
}
