package setup

import (
	"context"
	"database/sql"
	"time"
)

func ClearDB(DB *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		TRUNCATE TABLE users RESTART IDENTITY CASCADE;
	`
	_, err := DB.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
