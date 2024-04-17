package service

import (
	"context"
	"idon.com/models"
)

type PostService interface {
	AddPost(ctx context.Context, post models.Post) error
	UpdatePost(ctx context.Context, post models.Post) error
	DeletePost(ctx context.Context, id uint64) error

	GetPost(ctx context.Context, id uint64) (models.Post, error)
	GetPosts(ctx context.Context) (models.Posts, error)
}
