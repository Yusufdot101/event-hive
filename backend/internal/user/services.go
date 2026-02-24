package user

import (
	"errors"

	"github.com/Yusufdot101/eventhive/internal/customerrors"
	"github.com/go-playground/validator/v10"
)

type UserService struct {
	Repo *repository
}

func NewUserService(repo *repository) *UserService {
	return &UserService{Repo: repo}
}

func (us *UserService) RegisterUser(name, email, passwordPlaintext string) (*user, error) {
	u := &user{
		Name:  name,
		Email: email,
	}

	v := validator.New(validator.WithRequiredStructEnabled())
	err := validateUser(v, u)
	if err != nil {
		return nil, err
	}

	err = u.password.set(passwordPlaintext)
	if err != nil {
		return nil, err
	}

	return u, us.Repo.insert(u)
}

func (us *UserService) GetUserByEmailAndPassword(email, password string) (*user, error) {
	u, err := us.Repo.getByEmail(email)
	if err != nil {
		if errors.Is(err, customerrors.ErrNoRecord) {
			return nil, customerrors.ErrInvalidCredentials
		}
		return nil, err
	}

	matches, err := u.password.matches(password)
	if err != nil {
		return nil, err
	}

	if !matches {
		return nil, customerrors.ErrInvalidCredentials
	}

	return u, nil
}

func (us *UserService) GetUsersByIDs(userIDs []string) ([]*user, error) {
	return us.Repo.getManyByIDs(userIDs)
}
