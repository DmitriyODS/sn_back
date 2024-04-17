package app

import (
	"context"
	"idon.com/models"
)

func (a *App) AddLike(ctx context.Context, like models.Like) error {
	return a.reps.InsertLike(ctx, like)
}

func (a *App) DeleteLike(ctx context.Context, like models.Like) error {
	return a.reps.DeleteLike(ctx, like)
}

func (a *App) ToggleLike(ctx context.Context, like models.Like) error {
	return a.reps.ToggleLike(ctx, like)
}
