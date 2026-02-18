package event

import (
	"testing"
	"time"

	"github.com/Yusufdot101/eventhive/internal/config"
	"github.com/Yusufdot101/eventhive/internal/setup"
	"github.com/Yusufdot101/eventhive/internal/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	config.SetupVars()

	DB, err := config.OpenDB(config.TestDSN)
	if err != nil {
		t.Fatalf("unexpected error opening DB: %v", err)
	}

	err = setup.ClearDB(DB)
	if err != nil {
		t.Fatalf("unexpected error cleaning DB: %v", err)
	}

	userSvc := user.NewUserService(user.NewRepository(DB))
	u, err := userSvc.RegisterUser("yusuf", "ym@gmail.com", "12345678")
	if err != nil {
		t.Fatalf("unexpected error registering user: %v", err)
	}

	e := &event{
		creatorID:   u.ID,
		startsAt:    time.Now(),
		endsAt:      time.Now().Add(3 * 24 * time.Hour),
		title:       "test event",
		description: "this is a test event. please signup for it anyways",
		latitude:    0,
		longitude:   0,
	}

	repo := newRepository(DB)
	assert.NoError(t, repo.insert(e))
	assert.NoError(t, uuid.Validate(e.id))
}
