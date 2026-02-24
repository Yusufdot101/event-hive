package attendance

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/Yusufdot101/eventhive/internal/customerrors"
)

type repository struct {
	DB *sql.DB
}

func NewRepository(DB *sql.DB) *repository {
	return &repository{
		DB: DB,
	}
}

func (r *repository) insert(ea *eventAttendee) error {
	query := `
		INSERT INTO event_attendees
		(event_id, user_id)
		VALUES
		($1, $2)
		RETURNING created_at
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.DB.QueryRowContext(ctx, query, ea.EventID, ea.UserID).Scan(
		&ea.CreatedAt,
	)
}

func (r *repository) delete(ea *eventAttendee) error {
	query := `
		DELETE FROM event_attendees
		WHERE event_id = $1 
			AND user_id = $2
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := r.DB.ExecContext(ctx, query, ea.EventID, ea.UserID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return customerrors.ErrNoRecord
	}
	return nil
}

func (r *repository) get(ea *eventAttendee) (*eventAttendee, error) {
	query := `
		SELECT event_id, user_id FROM event_attendees
		WHERE event_id = $1 
			AND user_id = $2
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.DB.QueryRowContext(ctx, query, ea.EventID, ea.UserID).Scan(
		&ea.EventID,
		&ea.UserID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customerrors.ErrNoRecord
		}
		return nil, err
	}

	return ea, nil
}

func (r *repository) getManyByEventID(eventID string) ([]*eventAttendee, error) {
	query := `
		SELECT created_at, event_id, user_id FROM event_attendees
		WHERE event_id = $1 
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.DB.QueryContext(ctx, query, eventID)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("close rows error: ", err)
		}
	}()

	eventAttendees := []*eventAttendee{}
	for rows.Next() {
		ea := &eventAttendee{}
		err := rows.Scan(
			&ea.CreatedAt,
			&ea.EventID,
			&ea.UserID,
		)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, customerrors.ErrNoRecord
			}
			return nil, err
		}

		eventAttendees = append(eventAttendees, ea)
	}

	return eventAttendees, nil
}
