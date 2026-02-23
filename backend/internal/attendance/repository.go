package attendance

import (
	"context"
	"database/sql"
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

	return r.DB.QueryRowContext(ctx, query, ea.eventID, ea.userID).Scan(
		&ea.createdAt,
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

	res, err := r.DB.ExecContext(ctx, query, ea.eventID, ea.userID)
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
