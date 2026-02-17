package token

import (
	"time"

	"github.com/Yusufdot101/eventhive/internal/config"
	"github.com/Yusufdot101/eventhive/internal/customerrors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (ts *TokenService) GenerateRefreshToken(tokenUse tokenUse, userID string) (*token, error) {
	lifetime, err := time.ParseDuration(config.RefreshTokenLifetime)
	if err != nil {
		return nil, err
	}

	tk := &token{
		UserID:      userID,
		TokenString: uuid.New().String(),
		TokenUse:    tokenUse,
		ExpiresAt:   time.Now().Add(lifetime),
	}

	err = ts.repo.insert(tk)
	if err != nil {
		return nil, err
	}
	return tk, nil
}

func (ts *TokenService) GenerateJWT(tokenUse tokenUse, userID string) (*token, error) {
	lifetime, err := time.ParseDuration(config.JWTLifetime)
	if err != nil {
		return nil, err
	}

	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(lifetime).Unix(),
	})

	tokenString, err := tk.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return nil, err
	}

	return &token{
		UserID:      userID,
		TokenString: tokenString,
		TokenUse:    tokenUse,
		ExpiresAt:   time.Now().Add(lifetime),
	}, nil
}

func (ts *TokenService) GetTokenByStringAndUse(tokenString string, tokenUse tokenUse) (*token, error) {
	return ts.repo.getByStringAndUse(tokenString, tokenUse)
}

func (ts *TokenService) DeleteTokenByStringAndUse(tokenString string, tokenUse tokenUse) error {
	err := uuid.Validate(tokenString)
	if err != nil {
		return customerrors.ErrInvalidRefreshToken
	}

	return ts.repo.deleteByStringAndUse(tokenString, tokenUse)
}

func (ts *TokenService) ValidateJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, customerrors.ErrInvalidAccessToken
		}
		return []byte(config.JWTSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, customerrors.ErrInvalidAccessToken
	}

	return token, nil
}
