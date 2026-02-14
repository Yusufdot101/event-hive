package auth

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Yusufdot101/eventhive/internal/config"
	"github.com/Yusufdot101/eventhive/internal/setup"
	"github.com/Yusufdot101/eventhive/internal/token"
	"github.com/Yusufdot101/eventhive/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSignupHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	config.SetupVars()

	DB, err := config.OpenDB(config.TestDSN)
	if err != nil {
		log.Fatalf("unexpected error opening DB: %v", err)
	}
	err = setup.ClearDB(DB)
	if err != nil {
		log.Fatalf("unexpected error clearing DB: %v", err)
	}

	repo := user.NewRepository(DB)
	h := newHandler(user.NewUserService(repo), token.NewTokenService(token.NewRepository(DB)))

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

	h.signup(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "user created successfully")
}
