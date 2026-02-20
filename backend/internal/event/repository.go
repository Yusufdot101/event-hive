package event

import (
	"context"
	"database/sql"
	"time"
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
		e.startsAt,
		e.endsAt,
		e.creatorID,
		e.title,
		e.description,
		e.latitude,
		e.longitude,
		e.address,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.DB.QueryRowContext(ctx, query, values...).Scan(
		&e.id,
		&e.createdAt,
		&e.lastUpdatedAt,
	)
}
