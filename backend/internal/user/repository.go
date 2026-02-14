package user

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Yusufdot101/eventhive/internal/customerrors"
	_ "github.com/lib/pq"
)

type repository struct {
	DB *sql.DB
}

func NewRepository(DB *sql.DB) *repository {
	return &repository{
		DB: DB,
	}
}

func (r *repository) insert(u *user) error {
	query := `
		INSERT into users (name, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	values := []any{
		u.name,
		u.email,
		u.password.hash,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.DB.QueryRowContext(ctx, query, values...).Scan(
		&u.ID,
	)
	if err != nil {
		if err.Error() == `pq: duplicate key value violates unique constraint "users_email_key" (23505)` {
			return customerrors.ErrDuplicateEmail
		}
		return err
	}

	return nil
}

func (r *repository) getByEmail(email string) (*user, error) {
	query := `
		SELECT id, created_at, last_updated_at, name, email, password_hash FROM users
		WHERE email = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	u := &user{}
	err := r.DB.QueryRowContext(ctx, query, email).Scan(
		&u.ID,
		&u.createdAt,
		&u.lastUpdatedAt,
		&u.name,
		&u.email,
		&u.password.hash,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customerrors.ErrNoRecord
		}
		return nil, err
	}

	return u, nil
}
