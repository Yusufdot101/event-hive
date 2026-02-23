package event

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/Yusufdot101/eventhive/internal/customerrors"
	"github.com/lib/pq"
)

type repository struct {
	DB *sql.DB
}

func newRepository(DB *sql.DB) *repository {
	return &repository{
		DB: DB,
	}
}

func (r *repository) insert(e *event) error {
	query := `
		INSERT INTO events
		(starts_at, ends_at, creator_id, title, description, latitude, longitude, address)
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, last_updated_at
	`

	values := []any{
		e.StartsAt,
		e.EndsAt,
		e.CreatorID,
		e.Title,
		e.Description,
		e.Latitude,
		e.Longitude,
		e.Address,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.DB.QueryRowContext(ctx, query, values...).Scan(
		&e.ID,
		&e.CreatedAt,
		&e.LastUpdatedAt,
	)
}

func (r *repository) getMany() ([]*event, error) {
	query := `
		SELECT 
			id, created_at, starts_at, ends_at, last_updated_at, 
			creator_id, title, description, latitude, longitude, address
		FROM events
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	events := []*event{}
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("close rows error: ", err)
		}
	}()

	for rows.Next() {
		event := &event{}
		err = rows.Scan(
			&event.ID,
			&event.CreatedAt,
			&event.StartsAt,
			&event.EndsAt,
			&event.LastUpdatedAt,
			&event.CreatorID,
			&event.Title,
			&event.Description,
			&event.Latitude,
			&event.Longitude,
			&event.Address,
		)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func (r *repository) getByID(ID string) (*event, error) {
	query := `
		SELECT 
			id, created_at, starts_at, ends_at, last_updated_at, 
			creator_id, title, description, latitude, longitude, address
		FROM events
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	e := &event{}
	err := r.DB.QueryRowContext(ctx, query, ID).Scan(
		&e.ID,
		&e.CreatedAt,
		&e.StartsAt,
		&e.EndsAt,
		&e.LastUpdatedAt,
		&e.CreatorID,
		&e.Title,
		&e.Description,
		&e.Latitude,
		&e.Longitude,
		&e.Address,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customerrors.ErrNoRecord
		}
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			return nil, customerrors.ErrInvalidID
		}
		return nil, err
	}

	return e, nil
}
