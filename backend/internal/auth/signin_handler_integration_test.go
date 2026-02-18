package auth

import (
	"fmt"
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

func TestSigninHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	config.SetupVars()
	DB, err := config.OpenDB(config.TestDSN)
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
	err = setup.ClearDB(DB)
	if err != nil {
		log.Fatalf("unexpected error clearing DB: %v", err)
	}

	repo := user.NewRepository(DB)
	h := NewHandler(user.NewUserService(repo), token.NewTokenService(token.NewRepository(DB)))

	// register user first
	name := "yusuf"
	email := "ym@gmail.com"
	password := "12345678"

	_, err = h.userService.RegisterUser(name, email, password)
	if err != nil {
		log.Fatalf("unexpected error registering user: %v", err)
	}

	// signin user
	body := fmt.Sprintf(`{
		"name": "%s",
		"email": "%s",
		"password": "%s"
	}`, name, email, password)

	req := httptest.NewRequest(http.MethodPut, "/auth/signin", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	h.signin(c)

	assert.Equal(t, http.StatusOK, w.Code)
}
