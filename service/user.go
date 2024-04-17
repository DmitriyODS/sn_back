package service

import (
	"context"
	"errors"
	"idon.com/models"
)

var (
	ErrIncorrectLoginPass = errors.New("неверный логин, или пароль")
)

type AuthService interface {
	Login(ctx context.Context, user models.User) (models.User, error)
	Register(ctx context.Context, user models.User) (models.User, error)
}
