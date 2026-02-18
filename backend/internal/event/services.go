package event

import (
	"time"
)

func (svc *service) newEvent(creatorID string, startsAt, endsAt time.Time, title, description string, latitude, longitude float64) error {
	e := &event{
		creatorID:   creatorID,
		startsAt:    startsAt,
		endsAt:      endsAt,
		title:       title,
		description: description,
		latitude:    latitude,
		longitude:   longitude,
	}

	err := validateEvent(e)
	if err != nil {
		return err
	}
	return svc.repo.insert(e)
}
