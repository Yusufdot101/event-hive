package event

import (
	"testing"
	"time"

	"github.com/Yusufdot101/eventhive/internal/config"
	"github.com/Yusufdot101/eventhive/internal/customerrors"
	"github.com/Yusufdot101/eventhive/internal/setup"
	"github.com/Yusufdot101/eventhive/internal/user"
	"github.com/stretchr/testify/assert"
)

func TestNewEvent(t *testing.T) {
	config.SetupVars()
	DB, err := config.OpenDB(config.TestDSN)
	if err != nil {
		t.Fatalf("unexpected error opening DB: %v", err)
	}

	err = setup.ClearDB(DB)
	if err != nil {
		t.Fatalf("unexpected error clearing DB: %v", err)
	}

	repo := newRepository(DB)
	svc := newService(repo)

	// register user
	userSvc := user.NewUserService(user.NewRepository(DB))
	u, err := userSvc.RegisterUser("yusuf", "ym@gmail.com", "12345678")
	if err != nil {
		t.Fatalf("unexpected error registering user: %v", err)
	}

	_, err = svc.newEvent(u.ID, time.Now(), time.Now().Add(3*24*time.Hour), "test event", "test event description", 0, 0, "Test Address")
	if err != nil {
		t.Fatalf("unexpected error creating event: %v", err)
	}
}

func TestNewEventInvalid(t *testing.T) {
	config.SetupVars()
	DB, err := config.OpenDB(config.TestDSN)
	if err != nil {
		t.Fatalf("unexpected error opening DB: %v", err)
	}

	err = setup.ClearDB(DB)
	if err != nil {
		t.Fatalf("unexpected error clearing DB: %v", err)
	}

	repo := newRepository(DB)
	svc := newService(repo)

	// register user
	userSvc := user.NewUserService(user.NewRepository(DB))
	u, err := userSvc.RegisterUser("yusuf", "ym@gmail.com", "12345678")
	if err != nil {
		t.Fatalf("unexpected error registering user: %v", err)
	}

	startDate := time.Now()
	endDate := startDate.Add(-24 * time.Hour) // end date a day before start date
	_, err = svc.newEvent(u.ID, startDate, endDate, "test event", "test event description", 0, 0, "Test Address")
	assert.Equal(t, err.Error(), customerrors.ErrInvalidDates.Error())
}
