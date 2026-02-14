package token

import (
	"context"
	"database/sql"
	"errors"
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

func (r *repository) insert(tk *token) error {
	query := `
		INSERT INTO tokens (user_id, expires_at, token_string, token_use)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`

	values := []any{
		tk.UserID,
		tk.ExpiresAt,
		tk.TokenString,
		tk.TokenUse,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.DB.QueryRowContext(ctx, query, values...).Scan(
		&tk.ID,
		&tk.CreatedAt,
	)
}

func (r *repository) getByStringAndUse(tokenString string, tokenUse tokenUse) (*token, error) {
	query := `
		SELECT id, user_id, created_at, expires_at, token_string, token_use
		FROM tokens
		WHERE 
			token_string = $1
			AND token_use = $2 
			AND expires_at > NOW()
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tk := &token{}
	err := r.DB.QueryRowContext(ctx, query, tokenString, tokenUse).Scan(
		&tk.ID,
		&tk.UserID,
		&tk.CreatedAt,
		&tk.ExpiresAt,
		&tk.TokenString,
		&tk.TokenUse,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customerrors.ErrInvalidRefreshToken
		}
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			return nil, customerrors.ErrInvalidRefreshToken
		}
		return nil, err
	}

	return tk, nil
}
