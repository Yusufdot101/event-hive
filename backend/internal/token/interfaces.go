package token

import "time"

type (
	tokenUse string
)

var (
	RefreshToken tokenUse = "refresh"
	AccessToken  tokenUse = "access"
)

type token struct {
	ID          string
	UserID      string
	CreatedAt   time.Time
	ExpiresAt   time.Time
	TokenString string
	TokenUse    tokenUse
}

type TokenService struct {
	repo *repository
}

func NewTokenService(repo *repository) *TokenService {
	return &TokenService{
		repo: repo,
	}
}
