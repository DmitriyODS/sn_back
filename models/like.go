package models

import "context"

type Like struct {
	PostID uint64 `json:"post_id" db:"post_id"`
	UserID uint64 `json:"user_id" db:"user_id"`
}

type LikeRepository interface {
	InsertLike(ctx context.Context, like Like) error
	DeleteLike(ctx context.Context, like Like) error
}
