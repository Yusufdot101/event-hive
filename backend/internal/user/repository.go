package user

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
		u.Name,
		u.Email,
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
		&u.CreatedAt,
		&u.LastUpdatedAt,
		&u.Name,
		&u.Email,
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

func (r *repository) getManyByIDs(userIDs []string) ([]*user, error) {
	query := `
		SELECT id, created_at, last_updated_at, name, email FROM users
		WHERE id = ANY($1)
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.DB.QueryContext(ctx, query, pq.Array(userIDs))
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("close rows error: ", err)
		}
	}()

	users := []*user{}
	for rows.Next() {
		u := &user{}
		err = rows.Scan(
			&u.ID,
			&u.CreatedAt,
			&u.LastUpdatedAt,
			&u.Name,
			&u.Email,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (r *repository) getByID(userID string) (*user, error) {
	query := `
		SELECT id, created_at, last_updated_at, name, email FROM users
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	u := &user{}
	err := r.DB.QueryRowContext(ctx, query, userID).Scan(
		&u.ID,
		&u.CreatedAt,
		&u.LastUpdatedAt,
		&u.Name,
		&u.Email,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customerrors.ErrNoRecord
		}
		return nil, err
	}

	return u, nil
}
