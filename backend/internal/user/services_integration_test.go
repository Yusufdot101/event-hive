package user

import (
	"log"
	"testing"

	"github.com/Yusufdot101/eventhive/internal/config"
	"github.com/Yusufdot101/eventhive/internal/setup"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	name := "yusuf"
	email := "example@gmail.com"
	passwordPlaintext := "12345678"

	config.SetupVars()
	DB, err := config.OpenDB(config.TestDSN)
	if err != nil {
		log.Fatalf("unexpected error opening DB: %v", err)
	}

	err = setup.ClearDB(DB)
	if err != nil {
		log.Fatalf("unexpected error truncating DB: %v", err)
	}

	repo := &repository{
		DB: DB,
	}
	svc := NewUserService(repo)

	u, err := svc.RegisterUser(name, email, passwordPlaintext)
	if err != nil {
		log.Fatalf("unexpected error registering user: %v", err)
	}

	assert.Equal(t, u.name, name)
	assert.Equal(t, u.email, email)
	matches, err := u.password.matches(passwordPlaintext)
	if err != nil {
		log.Fatalf("unexpected error matching password: %v", err)
	}

	if !matches {
		t.Error("expected password to match")
	}
}
