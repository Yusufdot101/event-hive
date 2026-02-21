package event

import (
	"time"
)

func (svc *service) newEvent(CreatorID string, startsAt, endsAt time.Time, title, description string, latitude, longitude float64, address string) error {
	e := &event{
		CreatorID:   CreatorID,
		StartsAt:    startsAt,
		EndsAt:      endsAt,
		Title:       title,
		Description: description,
		Latitude:    latitude,
		Longitude:   longitude,
		Address:     address,
	}

	err := validateEvent(e)
	if err != nil {
		return err
	}
	return svc.repo.insert(e)
}

func (svc *service) getMany() ([]*event, error) {
	return svc.repo.getMany()
}
