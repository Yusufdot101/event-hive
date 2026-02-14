package token

import (
	"log"
	"testing"

	"github.com/Yusufdot101/eventhive/internal/config"
	"github.com/Yusufdot101/eventhive/internal/setup"
	"github.com/Yusufdot101/eventhive/internal/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGenerateRefreshToken(t *testing.T) {
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
	svc := NewTokenService(NewRepository(DB))
	tk, err := svc.GenerateRefreshToken(RefreshToken, u.ID)
	if err != nil {
		t.Fatalf("unexpected error inserting token: %v", err)
	}

	// assert the values
	assert.Equal(t, tk.UserID, u.ID)
	assert.Equal(t, tk.TokenUse, RefreshToken)
}

func TestGenerateJWT(t *testing.T) {
	// setup the DB
	config.SetupVars()
	DB, err := config.OpenDB(config.TestDSN)
	if err != nil {
		t.Fatalf("unexpected error opening DB: %v", err)
	}

	svc := NewTokenService(NewRepository(DB))
	userID := uuid.New().String()
	tk, err := svc.GenerateJWT(AccessToken, userID)
	if err != nil {
		t.Fatalf("unexpected error inserting token: %v", err)
	}

	// assert the values
	assert.Equal(t, tk.UserID, userID)
	assert.Equal(t, tk.TokenUse, AccessToken)
}
