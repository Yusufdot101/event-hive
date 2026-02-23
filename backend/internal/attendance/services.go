package attendance

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
