package customerrors

import "errors"

var (
	ErrDuplicateEmail      = errors.New("duplicate email")
	ErrNoRecord            = errors.New("record not found")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrInvalidRefreshToken = errors.New("invalid or expired refresh token")
)
