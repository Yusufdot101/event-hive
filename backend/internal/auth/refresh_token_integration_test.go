package auth

import (
	"fmt"
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

func TestRefreshTokenHandler(t *testing.T) {
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

	h := newHandler(user.NewUserService(
		user.NewRepository(DB)),
		token.NewTokenService(token.NewRepository(DB)),
	)

	// register user
	name := "yusuf"
	email := "example@gmail.com"
	passwordPlaintext := "12345678"

	body := fmt.Sprintf(`{
			"name": "%s",
			"email": "%s",
			"password": "%s"
	}`, name, email, passwordPlaintext)

	req := httptest.NewRequest(http.MethodPost, "/auth/signup", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	h.signup(ctx)

	assert.Equal(t, http.StatusCreated, w.Code)
	// requst refresh token
	// get the refresh token from the cookie
	cookie := w.Header().Get("Set-cookie")
	refreshToken := strings.Split(strings.Split(cookie, " ")[0], "=")[1]
	refreshToken = strings.Split(refreshToken, ";")[0] // remove the trailing ';'

	// use the refresh token to request a new access token
	req = httptest.NewRequest(http.MethodPut, "/auth/refreshtoken", nil)
	req.Header.Set("Cookie", fmt.Sprintf("refresh_token=%s", refreshToken))

	w = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(w)
	ctx.Request = req

	h.refreshToken(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")
}
