package config

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DSN                  string
	TestDSN              string
	RefreshTokenLifetime string
	JWTLifetime          string
	JWTSecret            string
	AllowedOrigins       []string
	AllowedMethods       []string
	AllowedHeaders       []string
	SecureCookie         bool
)

func SetupVars() {
	usr, _ := user.Current()
	home := usr.HomeDir
	loadEnv(home + "/Documents/projects/event-hive/backend/internal/config/.env")

	DSN = os.Getenv("DSN")
	TestDSN = os.Getenv("TEST_DSN")
	RefreshTokenLifetime = os.Getenv("REFRESH_TOKEN_LIFETIME")
	JWTLifetime = os.Getenv("JWT_LIFETIME")
	JWTSecret = os.Getenv("JWT_SECRET")

	AllowedOrigins = strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
	AllowedMethods = strings.Split(os.Getenv("ALLOWED_METHODS"), ",")
	AllowedHeaders = strings.Split(os.Getenv("ALLOWED_HEADERS"), ",")

	SecureCookie = os.Getenv("SECURE_COOKIE") != "false"
}

func loadEnv(envFileName string) {
	if err := godotenv.Load(envFileName); err != nil {
		log.Panicf("could not load .env file: %v", err)
	}
}

func OpenDB(DSN string) (*sql.DB, error) {
	DB, err := sql.Open("postgres", DSN)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = DB.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return DB, nil
}
