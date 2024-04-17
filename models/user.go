package models

import (
	"context"
	"errors"
)

type User struct {
	ID          uint64 `json:"id"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	AccessToken string `json:"access_token"`
}

type AuthRepository interface {
	SelectUser(ctx context.Context, login User) (uint64, error)
}

func (l User) Validate() error {
	if l.Login == "" {
		return errors.New("логин не может быть пустым")
	}

	if l.Password == "" {
		return errors.New("пароль не может быть пустым")
	}

	return nil
}
