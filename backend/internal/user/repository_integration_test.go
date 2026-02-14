package user

import (
	"log"
	"testing"

	"github.com/Yusufdot101/eventhive/internal/config"
	"github.com/Yusufdot101/eventhive/internal/setup"
)

func TestInsert(t *testing.T) {
	u := &user{
		ID:    "bb1ca442-5594-4106-aba1-34a40d89b9dc",
		name:  "yusuf",
		email: "example@gmail.com",
	}
	_ = u.password.set("12345678")

	config.SetupVars()
	DB, err := config.OpenDB(config.TestDSN)
	if err != nil {
		t.Fatalf("unexpected error opening DB: %v", err)
	}

	err = setup.ClearDB(DB)
	if err != nil {
		log.Fatalf("unexpected error truncating DB: %v", err)
	}

	repo := NewRepository(DB)

	err = repo.insert(u)
	if err != nil {
		t.Fatalf("unexpected error opening DB: %v", err)
	}
}

func TestGetByEmail(t *testing.T) {
	// instert user
	u := &user{
		ID:    "bb1ca442-5594-4106-aba1-34a40d89b9dc",
		name:  "yusuf",
		email: "example@gmail.com",
	}
	passwordPlaintext := "12345678"
	_ = u.password.set(passwordPlaintext)

	config.SetupVars()
	DB, err := config.OpenDB(config.TestDSN)
	if err != nil {
		t.Fatalf("unexpected error opening DB: %v", err)
	}

	err = setup.ClearDB(DB)
	if err != nil {
		log.Fatalf("unexpected error truncating DB: %v", err)
	}

	repo := NewRepository(DB)

	err = repo.insert(u)
	if err != nil {
		t.Fatalf("unexpected error opening DB: %v", err)
	}

	// fetch user
	gotUser, err := repo.getByEmail(u.email)
	if err != nil {
		t.Fatalf("unexpected error fetching user: %v", err)
	}

	// assert the values
	if gotUser.name != u.name {
		t.Errorf("expected gotUser.name=%s, got gotUser.name=%s", u.name, gotUser.name)
	}
	if gotUser.email != u.email {
		t.Errorf("expected gotUser.name=%s, got gotUser.name=%s", u.name, gotUser.name)
	}
	matches, err := gotUser.password.matches(passwordPlaintext)
	if gotUser.email != u.email {
		t.Fatalf("unexpected error matching password: %v", err)
	}
	if !matches {
		t.Error("expected gotUser.password to match set password")
	}
}
