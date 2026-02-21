package event

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Yusufdot101/eventhive/internal/auth"
	"github.com/Yusufdot101/eventhive/internal/config"
	"github.com/Yusufdot101/eventhive/internal/setup"
	"github.com/Yusufdot101/eventhive/internal/token"
	"github.com/Yusufdot101/eventhive/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetEventsHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	config.SetupVars()
	DB, err := config.OpenDB(config.TestDSN)
	if err != nil {
		t.Fatalf("unexpected error opening DB: %v", err)
	}

	err = setup.ClearDB(DB)
	if err != nil {
		t.Fatalf("unexpected error clearing DB: %v", err)
	}

	userSvc := user.NewUserService(user.NewRepository(DB))
	tokenSvc := token.NewTokenService(token.NewRepository(DB))
	authHandler := auth.NewHandler(userSvc, tokenSvc)

	// create user
	body := `{
		"name": "yusuf",
		"email": "example@gmail.com",
		"password": "12345678"
	}
	`
	req := httptest.NewRequest(http.MethodPost, "/auth/signup", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	authHandler.Signup(c)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response struct {
		Message     string `json:"message"`
		AccessToken string `json:"accessToken"`
	}

	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("unexpected error unmarshalling response: %v", err)
	}

	if response.AccessToken == "" {
		t.Fatal("expected accessToken to be in response")
	}

	// create event
	router := gin.New()
	group := router.Group("/")
	RegisterRoutes(DB, group)

	body = `{
		"title": "test event",
		"description": "test event description",
		"startsAt": "2026-02-18T11:35:20.123Z",
		"endsAt": "2026-03-18T11:35:20.123Z",
		"address": "test address"
	}`
	req = httptest.NewRequest(http.MethodPost, "/events", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+response.AccessToken)

	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// get the events
	h := newHandler(newService(newRepository(DB)))
	req = httptest.NewRequest(http.MethodGet, "/events", nil)
	w = httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	h.getEvents(ctx)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "events")
}
