package models

import "context"

type Like struct {
	PostID uint64
	UserID uint64
}

type LikeRepository interface {
	InsertLike(ctx context.Context, like Like) error
	DeleteLike(ctx context.Context, like Like) error
}
