package service

import (
	"context"
	"idon.com/models"
)

type LikeService interface {
	AddLike(ctx context.Context, like models.Like) error
	DeleteLike(ctx context.Context, like models.Like) error
}
