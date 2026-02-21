package customerrors

import "errors"

var (
	ErrDuplicateEmail      = errors.New("duplicate email")
	ErrNoRecord            = errors.New("record not found")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrInvalidRefreshToken = errors.New("invalid or expired refresh token")
	ErrInvalidAccessToken  = errors.New("invalid or expired access token")

	// ErrInvalidDates Doesn't include the createdAt because thats set by the server
	ErrInvalidDates    = errors.New("invalid event dates, either start time or end time")
	ErrInvalidInfo     = errors.New("invalid event info, either title, description or address")
	ErrInvalidLocation = errors.New("invalid event location, either longitude or latitude")
	ErrInvalidID       = errors.New("invalid id")
)
