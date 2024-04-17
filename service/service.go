package service

import "errors"

var (
	ErrService = errors.New("что-то пошло не так")
)

type Service interface {
	AuthService
	PostService
	LikeService
}
