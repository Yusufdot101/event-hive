package token

import (
	"log"
	"testing"
	"time"

	"github.com/Yusufdot101/eventhive/internal/config"
	"github.com/Yusufdot101/eventhive/internal/setup"
	"github.com/Yusufdot101/eventhive/internal/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFullFlow(t *testing.T) {
	// setup the DB
	config.SetupVars()
	DB, err := config.OpenDB(config.TestDSN)
	if err != nil {
		t.Fatalf("unexpected error opening DB: %v", err)
	}

	err = setup.ClearDB(DB)
	if err != nil {
		log.Fatalf("unexpected error truncating DB: %v", err)
	}

	// insert dummy user to avoid foreign key violations problems
	name := "yusuf"
	email := "example@gmail.com"
	passwordPlaintext := "12345678"
	userSvc := user.NewUserService(user.NewRepository(DB))
	u, err := userSvc.RegisterUser(name, email, passwordPlaintext)
	if err != nil {
		t.Fatalf("unexpected error inserting user: %v", err)
	}

	// create and insert the token
	repo := NewRepository(DB)
	tk := &token{
		UserID:      u.ID,
		ExpiresAt:   time.Now().Add(time.Hour),
		TokenUse:    RefreshToken,
		TokenString: uuid.New().String(),
	}

	err = repo.insert(tk)
	if err != nil {
		t.Fatalf("unexpected error inserting token: %v", err)
	}

	gotTk, err := repo.getByStringAndUse(tk.TokenString, tk.TokenUse)
	if err != nil {
		t.Fatalf("unexpected error fetching token: %v", err)
	}

	assert.Equal(t, gotTk.TokenString, tk.TokenString)
	assert.Equal(t, gotTk.UserID, tk.UserID)

	// normalize
	expected := tk.ExpiresAt.UTC()
	actual := gotTk.ExpiresAt.UTC()

	// Postgres stores microseconds, Go time.Time has nanoseconds.
	assert.WithinDuration(t, expected, actual, time.Microsecond)

	// delete
	err = repo.deleteByStringAndUse(gotTk.TokenString, gotTk.TokenUse)
	if err != nil {
		t.Fatalf("unexpected error deleting token: %v", err)
	}

	_, err = repo.getByStringAndUse(tk.TokenString, tk.TokenUse)
	if err == nil {
		t.Fatal("expected an error fetching token after deletion got none")
	}
}
