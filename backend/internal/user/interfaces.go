package user

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

const COST = 12

var errInvalidPassword = errors.New("invalid password")

type user struct {
	ID            string
	CreatedAt     time.Time
	LastUpdatedAt *time.Time
	Name          string   `validate:"gte=2"`
	Email         string   `validate:"email"`
	password      password `json:"-"`
}

type password struct {
	hash []byte
}

func (p *password) set(passwordPlaintext string) error {
	err := validatePasswordPlaintext(passwordPlaintext)
	if err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(passwordPlaintext), COST)
	if err != nil {
		return err
	}

	p.hash = hash
	return nil
}

func (p *password) matches(passwordPlaintext string) (bool, error) {
	err := validatePasswordPlaintext(passwordPlaintext)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword(p.hash, []byte(passwordPlaintext))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func validateUser(v *validator.Validate, u *user) error {
	err := v.Struct(u)
	if err != nil {
		// var invalidValidationError *validator.InvalidValidationError
		// if errors.As(err, &invalidValidationError) {
		// 	return err
		// } else {
		// 	return customError
		// }
		return err
	}

	return nil
}

func validatePasswordPlaintext(passwordPlaintext string) error {
	if len(passwordPlaintext) < 8 || len(passwordPlaintext) > 72 {
		return errInvalidPassword
	}
	return nil
}
