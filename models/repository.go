package models

import "errors"

var (
	ErrNoRows = errors.New("нет записей")
)

type Repository interface {
	AuthRepository
	PostRepository
	LikeRepository
}
