package models

type Repository interface {
	AuthRepository
	PostRepository
	LikeRepository
}
