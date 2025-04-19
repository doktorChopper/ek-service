package store

import (
	"errors"

	"github.com/doktorChopper/ek-service/internal/models"
)

var (
    ErrInvalidCredentials = errors.New("invalid credentials")
)

type AuthStore struct {
    userStore *UserStore
}

func NewAuthStore(u *UserStore) *AuthStore {
    return &AuthStore {
        userStore: u,
    }
}

// register new user

func (a *AuthStore) Register(u *models.User) error {

    _, err := a.userStore.Create(u)

    if err != nil {
        return err
    }

    return nil
}

func (a *AuthStore) Login(email, password string) (*models.User, error) {

    user, err := a.userStore.FindByEmail(email)
    if err != nil {
        return nil, ErrInvalidCredentials
    }

    if user.Password != password {
        return nil, ErrInvalidCredentials
    }

    return user, nil
}

